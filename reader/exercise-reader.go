package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
		b[x] = 'A'
	}
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

