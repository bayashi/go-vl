package vl

import (
	"os"
	"regexp"

	verticaltable "github.com/bayashi/go-verticaltable"
)

func Example_processLine() {
	v := &VL{
		Count: 0,
		Options: &Options{
			VtOpts: &verticaltable.VTOptions{
				HeaderFormat:  "********** %s **********",
				ShowCount:     false,
				CountFormat:   "%d. ",
				KvSeparator:   ": ",
				KeyAlignRight: true,
			},
			NoPager: true,
		},
	}

	v.processLine(os.Stdout, []byte("NAME                         READY     STATUS    RESTARTS   AGE"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-ntgdb   1/1       Running   0          9s"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-pc4j9   1/1       Running   0          9s"))
	// Output:
	// ********** 1 **********
	//     NAME: hello-web-4017757401-ntgdb
	//    READY: 1/1
	//   STATUS: Running
	// RESTARTS: 0
	//      AGE: 9s
	// ********** 2 **********
	//     NAME: hello-web-4017757401-pc4j9
	//    READY: 1/1
	//   STATUS: Running
	// RESTARTS: 0
	//      AGE: 9s
}

func Example_processLine_grep() {
	v := &VL{
		Count: 0,
		Options: &Options{
			VtOpts: &verticaltable.VTOptions{
				HeaderFormat:  "********** %s **********",
				ShowCount:     false,
				CountFormat:   "%d. ",
				KvSeparator:   ": ",
				KeyAlignRight: true,
			},
			NoPager: true,
			GrepRe: []*regexp.Regexp{
				regexp.MustCompile("pc4"),
				regexp.MustCompile("j9"),
			},
		},
	}

	v.processLine(os.Stdout, []byte("NAME                         READY     STATUS    RESTARTS   AGE"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-ntgdb   1/1       Running   0          9s"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-pc4j9   1/1       Running   0          9s"))
	// Output:
	// ********** 1 **********
	//     NAME: hello-web-4017757401-pc4j9
	//    READY: 1/1
	//   STATUS: Running
	// RESTARTS: 0
	//      AGE: 9s
}