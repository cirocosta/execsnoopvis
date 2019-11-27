#!/usr/bin/python

# trace.py  Traces the execution of processes.
#
# [execve]: https://elixir.bootlin.com/linux/v5.0/source/fs/exec.c#L1963
# [do_exit]: https://elixir.bootlin.com/linux/v5.0/source/kernel/exit.c#L773

from __future__ import division
from __future__ import print_function
from bcc import BPF
from pprint import pprint


program = BPF(
    text=r"""
#include <linux/sched.h>
#include <uapi/linux/ptrace.h>

BPF_PERF_OUTPUT(events);

enum event_type
{
        EVENT_START = 1,
        EVENT_FINISH,
};

struct data_t
{
        enum event_type type;

        char comm[TASK_COMM_LEN];
        int exitcode;
        u32 pid, ppid;
        u64 ts;
};

static inline void
__submit(struct pt_regs* ctx, struct data_t* data)
{
        struct task_struct* task;

        task = (struct task_struct*)bpf_get_current_task();

        data->ts = bpf_ktime_get_ns();
        data->pid = task->tgid;
        data->ppid = task->real_parent->tgid;
        bpf_get_current_comm(&data->comm, sizeof(data->comm));

        events.perf_submit(ctx, data, sizeof(*data));
}

static inline void
__submit_start(struct pt_regs* ctx)
{
        struct data_t data = {};

        data.type = EVENT_START;

        __submit(ctx, &data);
}

static inline void
__submit_finish(struct pt_regs* ctx, int code)
{
        struct data_t data = {};

        data.type = EVENT_FINISH;
        data.exitcode = code;

        __submit(ctx, &data);
}

int
kr__sys_execve(struct pt_regs* ctx,
               const char __user* filename,
               const char __user* const __user* __argv,
               const char __user* const __user* __envp)
{
        int ret = PT_REGS_RC(ctx);
        if (ret != 0) {
                return 0;
        }

        __submit_start(ctx);

        return 0;
}

int
k__do_exit(struct pt_regs* ctx, long code)
{
        __submit_finish(ctx, code);
        return 0;
}
"""
)

program.attach_kretprobe(
    event=program.get_syscall_fnname("execve"), fn_name="kr__sys_execve"
)

program.attach_kprobe(event="do_exit", fn_name="k__do_exit")


class EventType(object):
    EVENT_START = 1
    EVENT_FINISH = 2


procs = {}


def handle_events(cpu, data, size):
    event = program["events"].event(data)

    if event.type == EventType.EVENT_START:
        procs[event.pid] = event
        procs[event.pid].argv = get_cmdline(event.pid)

    elif event.type == EventType.EVENT_FINISH:
        if not event.pid in procs:
            return
        proc = procs[event.pid]
        elapsed = (event.ts - proc.ts) / (10 ** 9)

        print(
            "{:<16d} {:<16d} {:<16d} {:<16f} {}".format(
                proc.pid, proc.ppid, proc.exitcode, elapsed, " ".join(proc.argv),
            )
        )
        del procs[event.pid]



def get_cmdline(pid):
    try:
        with open("/proc/%d/cmdline" % pid) as f:
            return f.read().split('\0')
    except IOError:
        pass
    return []


program["events"].open_perf_buffer(handle_events)

print(
    "{:<16} {:<16} {:<16} {:<16} {:<16} ".format(
        "PID", "PPID", "CODE", "TIME(s)", "COMM"
    )
)

while 1:
    try:
        program.perf_buffer_poll()
    except KeyboardInterrupt:
        exit()
