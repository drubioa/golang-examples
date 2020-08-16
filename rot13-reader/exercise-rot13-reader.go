package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isUpperCase(b byte) bool{
	return int(b) >= 65 && int(b) <= 90
}

func isLowerCase(b byte) bool{
	return int(b) >= 97 && int(b) <= 122
}

func rot13(b byte) byte {

	if isUpperCase(b) {
		n := int(b) + 13
		if n > 90 {
			return byte(n - 25)
		} else {
			return byte(n)
		}
	}
	if isLowerCase(b) {
		n := int(b) + 13
		if n > 122 {
			return byte(n - 25)
		} else {
			return byte(n)
		}

	}
	return b
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

