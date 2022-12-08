package main

import "io"

type Trimmer struct {
	Input io.Reader
}

func NewTrimmer(in io.Reader) *Trimmer {
	return &Trimmer{
		Input: in,
	}
}

var _ io.Reader = (*Trimmer)(nil)

func (t *Trimmer) Read(p []byte) (n int, err error) {
	n, err = t.Input.Read(p)
	if err != nil {
		return n, err
	}

	outIdx := 0
	for _, c := range p[:n] {
		if isHexChar(c) {
			p[outIdx] = c
			outIdx++
		}
	}

	return outIdx, nil
}

func isHexChar(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'A' && c <= 'F' {
		return true
	}
	if c >= 'a' && c <= 'f' {
		return true
	}

	return false
}
