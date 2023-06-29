package main

import (
	"fmt"
	"os"
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
			NoPager: o.noPager,
		},
	}

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	out, closer := Pager(v.Options.NoPager)
	defer closer()

	v.Process(out)

	return nil
}
