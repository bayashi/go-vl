# go-vl

<a href="https://github.com/bayashi/go-vl/blob/main/LICENSE" title="go-vl License"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png" alt="MIT License"></a>
<a href="https://github.com/bayashi/go-vl/actions" title="go-vl CI"><img src="https://github.com/bayashi/go-vl/workflows/main/badge.svg" alt="go-vl CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/go-vl" title="go-vl report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/go-vl" alt="go-vl report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/go-vl" title="Go go-vl package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/go-vl.svg" alt="Go Reference: go-vl"></a>

`go-vl` provides `vl` filter to make CUI table vertical.

## Usage

For example, this is output of `kubectl get pods`.

```sh
$ kubectl get pods
NAME                         READY     STATUS    RESTARTS   AGE
hello-web-4017757401-ntgdb   1/1       Running   0          9s
hello-web-4017757401-pc4j9   1/1       Running   0          9s
```

The `vl` filter makes it vertical like below.

```sh
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

## Installation

```cmd
go install github.com/bayashi/go-vl@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
