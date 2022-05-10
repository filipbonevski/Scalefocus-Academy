package main

import (
	"io"
	"os"
	"strings"
)

//TASK 1

type ReverseStringReader struct {
	input string
}

func (r *ReverseStringReader) Read(p []byte) (int, error) {
	var b strings.Builder
	b.Grow(len(r.input))
	for i := len(r.input) - 1; i >= 0; i-- {
		b.WriteByte(r.input[i])
	}

	reversedString := b.String()

	n := copy(p, []byte(reversedString))

	return n, io.EOF
}

func NewReverseStringReader(input string) *ReverseStringReader {
	return &ReverseStringReader{input}
}

func main() {
	io.Copy(os.Stdout, NewReverseStringReader("palindrom"))
}
