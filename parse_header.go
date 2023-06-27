package vl

import (
	"regexp"
)

type Column struct {
	Label string
}

type Header struct {
	Columns []*Column
}

var splitter = regexp.MustCompile(`\s\s+`)

func ParseHeader(line []byte) *Header {
	labels := splitter.Split(string(line), -1)

	hs := &Header{}
	for _, label := range labels {
		c := &Column{
			Label: label,
		}
		hs.Columns = append(hs.Columns, c)
	}

	return hs
}
