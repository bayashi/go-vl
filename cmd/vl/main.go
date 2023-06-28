package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"

	verticaltable "github.com/bayashi/go-verticaltable"
	vl "github.com/bayashi/go-vl"
	"golang.org/x/term"
)

func main() {
	err := run()
	if err != nil {
		putErr(fmt.Sprintf("Err %s: %s", cmd, err.Error()))
		os.Exit(exitErr)
	}

	os.Exit(exitOK)
}

func run() error {
	o := parseArgs()
	v := &vl.VL{
		Count: 0,
		Options: &vl.Options{
			GrepRe: o.grepRe,
		},
	}

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	s := bufio.NewScanner(os.Stdin)

	out, closer := Pager(o)
	defer closer()

	for s.Scan() {
		line := s.Bytes()

		if len(line) == 0 {
			continue
		}

		if v.Count == 0 {
			v.Header = vl.ParseHeader(line)
		}

		if v.Count > 0 {
			if len(v.Options.GrepRe) > 0 && vl.IsFiltered(v, line) {
				continue
			}
			elements := vl.Process(v, line)
			vt := verticaltable.NewTable(out, vtOpts())
			vt.Header(strconv.Itoa(v.Count))
			for i, elem := range elements {
				vt.Row(v.Header.Columns[i].Label, elem)
			}
			vt.Render()
		}

		v.Count++
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
