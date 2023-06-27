package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"

	verticaltable "github.com/bayashi/go-verticaltable"
	"golang.org/x/term"
)

type VL struct {
	o      *options
	count  int
	header *header
}

func main() {
	err := run()
	if err != nil {
		putErr(fmt.Sprintf("Err %s: %s", cmd, err.Error()))
		os.Exit(exitErr)
	}

	os.Exit(exitOK)
}

func run() error {
	vl := &VL{
		o:     parseArgs(),
		count: 0,
	}

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	s := bufio.NewScanner(os.Stdin)

	out, closer := Pager(vl.o)
	defer closer()

	for s.Scan() {
		line := s.Bytes()

		if len(line) == 0 {
			continue
		}

		if vl.count == 0 {
			vl.header = parseHeader(line)
		}

		if vl.count > 0 {
			elements := process(vl, line)
			vt := verticaltable.NewTable(out, vtOpts())
			vt.Header(strconv.Itoa(vl.count))
			for i, elem := range elements {
				vt.Row(vl.header.columns[i].label, elem)
			}
			vt.Render()
		}

		vl.count++
	}

	return nil
}

func vtOpts() *verticaltable.VTOptions {
	return &verticaltable.VTOptions{
		HeaderFormat:  "********** %s **********",
		ShowCount:     false,
		CountFormat:   "%d. ",
		KvSeparator:   ": ",
		KeyAlignRight: true,
	}
}
