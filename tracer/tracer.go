package tracer

import "C"

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/iovisor/gobpf/bcc"
)

// Execution represents a complete command execution that started and finished
// while the tracer was active.
//
type Execution struct {
	Pid, Ppid uint32
	ExitCode  int
	Time      uint64
	Argv      []string

	Start, Finish uint64
}

// payload is the data communicated between the kernel and userspace through a
// perf map.
//
type payload struct {
	Type int32

	Comm      [16]byte
	ExitCode  int32
	Pid, Ppid uint32
	Ts        uint64
}

type EventType int32

const (
	EventStart EventType = iota + 1
	EventFinish
)

type Tracer struct {
	procs map[uint32]Execution
	mod   *bcc.Module
}

func New() (t Tracer) {
	return Tracer{
		mod:   bcc.NewModule(src, []string{}),
		procs: make(map[uint32]Execution, 0),
	}
}

func (t *Tracer) Close() {
	t.mod.Close()
}

func (t *Tracer) Run(ctx context.Context) (err error) {
	err = attachProbes(t.mod)
	if err != nil {
		err = fmt.Errorf("attaching probes: %w", err)
		return
	}

	err = t.loop(ctx)
	if err != nil {
		err = fmt.Errorf("event loop: %w", err)
		return
	}

	return
}

func (t *Tracer) loop(ctx context.Context) (err error) {
	var (
		evs    = bcc.NewTable(t.mod.TableId("events"), t.mod)
		evsC   = make(chan []byte, 100)
		execsC = make(chan Execution, 1)
	)

	perfMap, err := bcc.InitPerfMap(evs, evsC)
	if err != nil {
		err = fmt.Errorf("perf map: %w", err)
		return
	}

	defer perfMap.Stop()

	go func() {
		for data := range evsC {
			dispatch(data, t.start, t.finish, execsC)
		}
	}()

	perfMap.Start()

	fmt.Printf("%-16s %-16s %-16s %-16s %-16s \n",
		"PID", "PPID", "CODE", "TIME(s)", "ARGV")

	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		case execution := <-execsC:
			fmt.Printf("%-16d %-16d %-16d %-16d %-16s \n",
				execution.Pid,
				execution.Ppid,
				execution.ExitCode,
				0,
				strings.Join(execution.Argv, " "))
		}
	}

	return
}

func procArgv(pid uint32) (argv []string, err error) {
	f, err := os.Open("/proc/" + strconv.Itoa(int(pid)) + "/cmdline")
	if err != nil {
		return
	}

	defer f.Close()

	s, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	for _, b := range bytes.Split(s, []byte{'\x00'}) {
		argv = append(argv, string(b))
	}
	return
}

func (t *Tracer) start(p payload) (err error) {
	execution := Execution{
		Pid:   p.Pid,
		Ppid:  p.Ppid,
		Start: p.Ts,
	}

	execution.Argv, err = procArgv(p.Pid)
	if err != nil {
		return
	}

	t.procs[p.Pid] = execution
	return
}

func (t *Tracer) finish(p payload) (execution *Execution, err error) {
	e, found := t.procs[p.Pid]
	if !found {
		return
	}

	delete(t.procs, p.Pid)
	execution = &e
	return
}

func dispatch(data []byte,
	start func(payload) error,
	finish func(payload) (*Execution, error),
	results chan<- Execution,
) (err error) {
	var (
		p         payload
		execution *Execution
	)

	err = binary.Read(bytes.NewBuffer(data), bcc.GetHostByteOrder(), &p)
	if err != nil {
		err = fmt.Errorf("decoding received data: %w", err)
		return
	}

	switch EventType(p.Type) {
	case EventStart:
		err = start(p)
	case EventFinish:
		execution, err = finish(p)
		if err != nil {
			return
		}

		if execution == nil {
			return
		}

		results <- *execution
		return
	default:
		err = fmt.Errorf("unknown ev %d", p.Type)
	}

	return
}

func attachProbes(mod *bcc.Module) (err error) {
	err = attachStartProbe(mod)
	if err != nil {
		return
	}

	err = attachFinishProbe(mod)
	return
}

func attachStartProbe(mod *bcc.Module) (err error) {
	const prog = `kr__sys_execve`
	var sys = bcc.GetSyscallFnName("execve")

	probe, err := mod.LoadKprobe(prog)
	if err != nil {
		return
	}

	err = mod.AttachKretprobe(sys, probe, -1)
	return
}

func attachFinishProbe(mod *bcc.Module) (err error) {
	const (
		prog = `k__do_exit`
		sys  = `do_exit`
	)

	probe, err := mod.LoadKprobe(prog)
	if err != nil {
		return
	}

	err = mod.AttachKprobe(sys, probe, -1)
	return
}
