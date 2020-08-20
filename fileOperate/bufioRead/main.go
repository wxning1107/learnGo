package main

import (
	"bufio"
	"io"
	"os"
)

/**
当文件较小（KB 级别）时，ioutil > bufio > os。
当文件大小比较常规（MB 级别）时，三者差别不大，但 bufio 又是已经显现出来。
当文件较大（GB 级别）时，bufio > os > ioutil。
*/

func Read(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	buf := make([]byte, 4096)
	var nByte int
	reader := bufio.NewReader(file)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err.Error())
		}
		if n == 0 {
			break
		}
		nByte += n
	}
	return nByte
}

func main() {

}
