package vl

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"

	verticaltable "github.com/bayashi/go-verticaltable"
)

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

		if len(line) == 0 {
			continue
		}

		if v.Count == 0 {
			v.Header = v.ParseHeader(line)
		}

		if v.Count > 0 {
			if len(v.Options.GrepRe) > 0 && v.isFiltered(line) {
				continue
			}

			vt := verticaltable.NewTable(out, v.Options.VtOpts)
			vt.Header(strconv.Itoa(v.Count))
			for i, elem := range v.processLine(line) {
				if !v.Header.Columns[i].Show {
					continue
				}
				vt.Row(v.Header.Columns[i].Label, elem)
			}
			vt.Render()
		}

		v.Count++
	}
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

func (v *VL) processLine(origLine []byte) []string {
	return splitter.Split(string(origLine), len(v.Header.Columns))
}
