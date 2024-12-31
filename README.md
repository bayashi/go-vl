# go-vl

<a href="https://github.com/bayashi/go-vl/blob/main/LICENSE" title="go-vl License"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png" alt="MIT License"></a>
<a href="https://github.com/bayashi/go-vl/actions" title="go-vl CI"><img src="https://github.com/bayashi/go-vl/workflows/main/badge.svg" alt="go-vl CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/go-vl" title="go-vl report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/go-vl" alt="go-vl report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/go-vl" title="Go go-vl package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/go-vl.svg" alt="Go Reference: go-vl"></a>

`go-vl` provides `vl` filter to make CUI table vertical.

## Usage

For example, this is output of `kubectl get pods`.

```cmd
$ kubectl get pods
NAME                         READY     STATUS    RESTARTS   AGE
hello-web-4017757401-ntgdb   1/1       Running   0          9s
hello-web-4017757401-pc4j9   1/1       Running   0          9s
```

The `vl` filter makes it vertical like below.

```
$ kubectl get pods | vl
********** 1 ********************
    NAME: hello-web-4017757401-ntgdb
   READY: 1/1
  STATUS: Running
RESTARTS: 0
     AGE: 9s
********** 2 ********************
    NAME: hello-web-4017757401-pc4j9
   READY: 1/1
  STATUS: Running
RESTARTS: 0
     AGE: 9s
```

`--grep` option works.

```
$ kubectl get pods | vl --grep pc4j9
********** 1 **********
    NAME: hello-web-4017757401-pc4j9
   READY: 1/1
  STATUS: Running
RESTARTS: 0
     AGE: 9s
```

And `--ps` flag to parse lines separated by only 1 space between columns.

```
$ ps
  PID TTY          TIME CMD
12523 pts/4    00:00:00 bash
30087 pts/4    00:00:00 ps

$ ps | vl --ps
********** 1 **********
 PID: 12523
 TTY: pts/4
TIME: 00:00:00
 CMD: bash
********** 2 **********
 PID: 30087
 TTY: pts/4
TIME: 00:00:00
 CMD: ps
```

## Full Options

```
$ vl --help
Usage: some-command | vl OPTIONS
Options:
  -g, --grep stringArray   Grep condition to filter lines
  -h, --help               Show help (This message) and exit
  -l, --label string       Show only matching items of labels
      --no-pager           Output without pager
      --ps                 Parse lines separated by one space. i.e. 'ps' command
  -v, --version            Show version and build info and exit
```

## Installation

```cmd
go install github.com/bayashi/go-vl/cmd/vl@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
