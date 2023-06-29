package main

// Most code was copied from https://github.com/jackdoe/go-pager

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const (
	ENV_KEY_PAGER = "PAGER"
	ENV_VALUE_NO_PAGER = "NOPAGER"
)

func Pager(noPager bool) (io.Writer, func()) {
	if noPager {
		return os.Stdout, func() {}
	}

	p, err := pagerPath("less", "more", "lv", "cat")
	if err != nil {
		putErr(err.Error())
		os.Exit(exitErr)
	}

	if p != "" {
		signal.Ignore(syscall.SIGPIPE)

		c := exec.Command(p)
		r, w := io.Pipe()
		c.Stdin = r
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		ch := make(chan struct{})
		go func() {
			defer close(ch)
			err := c.Run()
			if err != nil {
				panic(err)
			}
			os.Exit(exitOK)
		}()

		return w, func() {
			w.Close()
			<-ch
		}
	}

	return os.Stdout, func() {}
}

func pagerPath(pagers ...string) (string, error) {
	p := os.Getenv(ENV_KEY_PAGER)
	if p != "" {
		if p == ENV_VALUE_NO_PAGER {
			return "", nil
		}

		exe, err := exec.LookPath(p)
		if err != nil {
			return "", fmt.Errorf("the value `%s` for ENV:%s is wrong: %w", p, ENV_KEY_PAGER, err)
		}
		return exe, nil
	}

	for _, x := range pagers {
		exe, err := exec.LookPath(x)
		if err != nil {
			return "", fmt.Errorf("PAGER:%s is wrong: %w", x, err)
		} else {
			return exe, nil
		}
	}

	return "", nil
}
