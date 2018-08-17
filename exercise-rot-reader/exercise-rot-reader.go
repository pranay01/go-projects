package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader)Read(b []byte) (int, error) {
	
	n, err := rot.r.Read(b)

	if err == nil {
		for i:=0;i<n;i++ {
			//fmt.Println(b[i])
			// Removing the base -> Adding 13
			// taking mod 26 for circular rotation -> adding back the base
			if b[i]>64 && b[i]<91 {
				b[i] = (b[i]-64+13)%26 + 64 /
			} 
			if b[i]>96 && b[i]<123 { 
				b[i] = (b[i]-96+13)%26 + 96
			}
			//fmt.Println(string(b[i]))
		}
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
