package main

import (
	"bytes"
	"fmt"
	"learnGoSource/utils"
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
	fmt.Println()
	src := []string{"b", "c"}
	dst := utils.AppendFromIndex(src, "a", 0)
	fmt.Println(dst)
}
