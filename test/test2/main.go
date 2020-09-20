package main

import (
	"bytes"
	"os"
)

func main() {
	s := "qingting"
	var b bytes.Buffer
	b.WriteString(s)
	_, err := b.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
}
