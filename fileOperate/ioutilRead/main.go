package main

import (
	"io/ioutil"
	"os"
)

// ioutil.ReadAll Read to mem once
func Read(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)

	return len(all)
}

func main() {

}
