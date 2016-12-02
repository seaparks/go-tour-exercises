package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error){
	int, err := rot.r.Read(b)
	for i := range b {
		if ( b[i] >= 'A' && b[i] < 'N' || b[i] >= 'a' && b[i] < 'n' ) {
			b[i] = b[i] + 13
		} else if ( b[i] >= 'N' && b[i] <= 'Z' || b[i] >= 'n' && b[i] <= 'z' ) {
			b[i] = b[i] - 13
		}
	}
	return int, err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

