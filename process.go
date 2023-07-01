package vl

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"

	verticaltable "github.com/bayashi/go-verticaltable"
)

type Column struct {
	Label string
	Show  bool
}

type Header struct {
	Columns []*Column
}

type Options struct {
	GrepRe  []*regexp.Regexp
	Labels  []string
	VtOpts  *verticaltable.VTOptions
	NoPager bool
}

type VL struct {
	Count   int
	Header  *Header
	Options *Options
}

func (v *VL) Process(out io.Writer) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Bytes()
		v.processLine(out, line)
	}
}

func (v *VL) processLine(out io.Writer, origLine []byte) {
	if len(origLine) == 0 {
		return
	}

	if v.Count == 0 {
		v.Header = v.ParseHeader(origLine)
	}

	if v.Count > 0 {
		if len(v.Options.GrepRe) > 0 && v.isFiltered(origLine) {
			return
		}

		vt := verticaltable.NewTable(out, v.Options.VtOpts)
		vt.Header(strconv.Itoa(v.Count))
		for i, elem := range v.parseLine(origLine) {
			if !v.Header.Columns[i].Show {
				return
			}
			vt.Row(v.Header.Columns[i].Label, elem)
		}
		vt.Render()
	}

	v.Count++
}

var splitter = regexp.MustCompile(`\s\s+`)

func (v *VL) ParseHeader(line []byte) *Header {
	labels := splitter.Split(string(line), -1)

	hs := &Header{}
	for _, label := range labels {
		c := &Column{
			Label: label,
			Show: isShownLabel(label, v.Options.Labels),
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

// true:filtering, false:show
func (v *VL) isFiltered(origLine []byte) bool {
	for _, r := range v.Options.GrepRe {
		if !r.Match(origLine) {
			return true
		}
	}

	return false
}

func (v *VL) parseLine(origLine []byte) []string {
	return splitter.Split(string(origLine), len(v.Header.Columns))
}
