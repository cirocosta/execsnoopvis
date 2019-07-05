package parser

import (
	"bufio"
	"io"
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

	if len(fields) != 5 {
		err = errors.Errorf("not enough fields")
		return
	}

	return
}
