// Code for https://tour.golang.org/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (rd MyReader) Read(b []byte) (int, error) {
	len := 20

	for i := 0; i < len; i++ {
		b[i] = 'A'
	}

	return len, nil
}

func main() {
	reader.Validate(MyReader{})
}
