package vl

import (
	"regexp"
)

type Column struct {
	Label string
	Show  bool
}

type Header struct {
	Columns []*Column
}

var splitter = regexp.MustCompile(`\s\s+`)

func ParseHeader(line []byte, o *Options) *Header {
	labels := splitter.Split(string(line), -1)

	hs := &Header{}
	for _, label := range labels {
		c := &Column{
			Label: label,
			Show: isShownLabel(label, o.Labels),
		}
		hs.Columns = append(hs.Columns, c)
	}

	return hs
}

func isShownLabel(label string, labels []string) bool {
	if len(labels) == 0 {
		return true
	}

	for _, l := range labels {
		if l == label {
			return true
		}
	}

	return false
}
