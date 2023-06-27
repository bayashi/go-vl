package main

func process(vl *VL, origLine []byte) ([]string) {
	return splitter.Split(string(origLine), len(vl.header.columns))
}
