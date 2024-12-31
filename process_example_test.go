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

func Example_processLine_label() {
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
			Labels: []string{
				"READY",
				"STATUS",
			},
		},
	}

	v.processLine(os.Stdout, []byte("NAME                         READY     STATUS    RESTARTS   AGE"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-ntgdb   1/1       Running   0          9s"))
	v.processLine(os.Stdout, []byte("hello-web-4017757401-pc4j9   1/1       Running   0          9s"))
	// Output:
	// ********** 1 **********
	//  READY: 1/1
	// STATUS: Running
	// ********** 2 **********
	//  READY: 1/1
	// STATUS: Running
}

func Example_processLine_ps() {
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
			PS:      true,
		},
	}

	v.processLine(os.Stdout, []byte("USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND"))
	v.processLine(os.Stdout, []byte("root         1  0.0  0.0    904   520 ?        Sl   07:06   0:00 /init"))
	v.processLine(os.Stdout, []byte("root       175  0.0  0.0  34840  2200 ?        Ss   07:06   0:00 nginx: master process /usr/local/nginx/sbin/nginx"))
	v.processLine(os.Stdout, []byte("user     10421  0.7  3.0 1941284 121992 pts/0  Sl+  09:32   0:56 /home/user/go/bin/gopls -mode=stdio"))
	// Output:
	// ********** 1 **********
	//    USER: root
	//     PID: 1
	//    %CPU: 0.0
	//    %MEM: 0.0
	//     VSZ: 904
	//     RSS: 520
	//     TTY: ?
	//    STAT: Sl
	//   START: 07:06
	//    TIME: 0:00
	// COMMAND: /init
	// ********** 2 **********
	//    USER: root
	//     PID: 175
	//    %CPU: 0.0
	//    %MEM: 0.0
	//     VSZ: 34840
	//     RSS: 2200
	//     TTY: ?
	//    STAT: Ss
	//   START: 07:06
	//    TIME: 0:00
	// COMMAND: nginx: master process /usr/local/nginx/sbin/nginx
	// ********** 3 **********
	//    USER: user
	//     PID: 10421
	//    %CPU: 0.7
	//    %MEM: 3.0
	//     VSZ: 1941284
	//     RSS: 121992
	//     TTY: pts/0
	//    STAT: Sl+
	//   START: 09:32
	//    TIME: 0:56
	// COMMAND: /home/user/go/bin/gopls -mode=stdio
}
