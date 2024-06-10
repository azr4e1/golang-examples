package main

import (
	"io"
	"os"
	"strings"
)

const (
	AlphabetLength = 26
	Rot13          = 13
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	var base byte
	for i, el := range b {
		switch {
		case el == ' ':
			continue
		case el >= 'A' && el <= 'A'+AlphabetLength:
			base = 'A'
		default:
			base = 'a'
		}

		b[i] = ((el-base)+Rot13)%AlphabetLength + base
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
