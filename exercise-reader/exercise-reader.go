package main

import "golang.org/x/tour/reader"

//Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader)Read([]byte) (int, error) {
	return 'A'
}

func main() {
	reader.Validate(MyReader{})
}
