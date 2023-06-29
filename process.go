package vl

import (
	"regexp"

	verticaltable "github.com/bayashi/go-verticaltable"
)

type Options struct {
	GrepRe []*regexp.Regexp
	Labels []string
	VtOpts *verticaltable.VTOptions
}

type VL struct {
	Count   int
	Header  *Header
	Options *Options
}

// true:filtering, false:show
func (v *VL) IsFiltered(origLine []byte) bool {
	for _, r := range v.Options.GrepRe {
		if !r.Match(origLine) {
			return true
		}
	}

	return false
}

func (v *VL) ProcessLine(origLine []byte) []string {
	return splitter.Split(string(origLine), len(v.Header.Columns))
}
