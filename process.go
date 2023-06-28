package vl

import "regexp"

type Options struct {
	GrepRe []*regexp.Regexp
}

type VL struct {
	Count   int
	Header  *Header
	Options *Options
}

// true:filtering, false:show
func IsFiltered(v *VL, origLine []byte) bool {
	for _, r := range v.Options.GrepRe {
		if !r.Match(origLine) {
			return true
		}
	}

	return false
}

func Process(v *VL, origLine []byte) []string {
	return splitter.Split(string(origLine), len(v.Header.Columns))
}
