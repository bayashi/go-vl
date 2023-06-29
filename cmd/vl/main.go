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
			Labels: o.labels,
			VtOpts: &verticaltable.VTOptions{
				HeaderFormat:  "********** %s **********",
				ShowCount:     false,
				CountFormat:   "%d. ",
				KvSeparator:   ": ",
				KeyAlignRight: true,
			},
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
			v.Header = v.ParseHeader(line)
		}

		if v.Count > 0 {
			if len(v.Options.GrepRe) > 0 && v.IsFiltered(line) {
				continue
			}

			vt := verticaltable.NewTable(out, v.Options.VtOpts)
			vt.Header(strconv.Itoa(v.Count))
			for i, elem := range v.ProcessLine(line) {
				if !v.Header.Columns[i].Show {
					continue
				}
				vt.Row(v.Header.Columns[i].Label, elem)
			}
			vt.Render()
		}

		v.Count++
	}

	return nil
}
