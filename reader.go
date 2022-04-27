package gosign

import (
	"bytes"
	"io"
)

type ContainsReader struct {
	reader   io.Reader
	contains bool
	subslice []byte
}

func NewContainsReader(reader io.Reader, subslice []byte) *ContainsReader {
	return &ContainsReader{reader: reader, subslice: subslice}
}

func (p *ContainsReader) Read(data []byte) (n int, err error) {
	if !p.contains {
		p.contains = bytes.Contains(data, p.subslice)
	}

	return p.reader.Read(data)
}

func (p *ContainsReader) Contains() bool {
	return p.contains
}
