package main

import "golang.org/x/tour/reader"

//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader)Read(b []byte) (int, error) {
	
	// Just fill whatever byte you get with 'A's
	length := len(b)
	for i:=0;i<length;i++ {
		b[i] = 'A'
	}
	return length, nil
}

func main() {
	reader.Validate(MyReader{})
}
