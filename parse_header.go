package main

import (
	"regexp"
)

type column struct {
	label string
}

type header struct {
	columns []*column
}

var splitter = regexp.MustCompile(`\s\s+`)

func parseHeader(line []byte) *header {
	labels := splitter.Split(string(line), -1)

	hs := &header{}
	for _, label := range labels {
		c := &column{
			label: label,
		}
		hs.columns = append(hs.columns, c)
	}

	return hs
}
