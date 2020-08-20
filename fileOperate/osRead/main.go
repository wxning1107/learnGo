package main

import (
	"io"
	"os"
)

func Read(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	buf := make([]byte, 4096)
	var nBytes int
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
		if n == 0 {
			break
		}
		nBytes += n
	}

	return nBytes
}

func main() {

}
