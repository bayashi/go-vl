package vl

type VL struct {
	Count  int
	Header *Header
}

func Process(v *VL, origLine []byte) ([]string) {
	return splitter.Split(string(origLine), len(v.Header.Columns))
}
