package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Decode(c byte) (byte, error) {
	if c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm' {
		return c + 13, nil
	} else if c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z' {
		return c - 13, nil
	}
	return c, nil
}

func (rd *rot13Reader) Read(b []byte) (int, error) {
	len, err := rd.r.Read(b)
	for i := 0; i < len; i++ {
		b[i], err = rot13Decode(b[i])
	}
	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
