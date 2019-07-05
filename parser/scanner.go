package parser

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrIsHeader = errors.Errorf("header detected")
)

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner(reader io.Reader) Scanner {
	return Scanner{
		scanner: bufio.NewScanner(reader),
	}
}

// ScanAll scans `Nodes` from start to end, consuming everything from the reader
// until it has been completely consumed.
//
func (p *Scanner) ScanAll() (nodes []*Node, err error) {
	var (
		node Node
		done bool
	)

	for {
		node, done, err = p.Scan()
		if err != nil {
			if err == ErrIsHeader {
				continue
			}

			err = errors.Wrapf(err, "failed to scan line")
			return
		}

		if done {
			return
		}

		nodes = append(nodes, &node)
	}

	return
}

// Scan cosumes a single line from the reader.
//
func (p *Scanner) Scan() (node Node, done bool, err error) {
	done = !p.scanner.Scan()
	if done {
		err = errors.Wrapf(p.scanner.Err(),
			"failed to scan")
		return
	}

	line := p.scanner.Text()

	node, err = parseLine(line)
	if err != nil {
		err = errors.Wrapf(err,
			"failed parsing the line `%s`", line)
		return
	}

	return
}

func parseLine(line string) (node Node, err error) {
	fields := strings.Fields(line)

	if len(fields) < 5 {
		err = errors.Errorf("not enough fields")
		return
	}

	if fields[0] == "PCOMM" && fields[1] == "PID" {
		err = ErrIsHeader
		return
	}

	node.Command = fields[0]
	node.Pid, err = strconv.ParseUint(fields[1], 10, 32)
	if err != nil {
		err = errors.Wrapf(err,
			"pid field `%s` is not an unsigned int", fields[1])
		return
	}

	node.Ppid, err = strconv.ParseUint(fields[2], 10, 32)
	if err != nil {
		err = errors.Wrapf(err,
			"ppid field `%s` is not an unsigned int", fields[2])
		return
	}

	node.ExitCode, err = strconv.Atoi(fields[3])
	if err != nil {
		err = errors.Wrapf(err,
			"ppid field `%s` is not an int", fields[3])
		return
	}

	node.Argv = fields[4:len(fields)]

	return
}
