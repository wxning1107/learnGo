package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

/**
ReadBytes对小文件处理效率最差
在处理大文件时，ReadLine和ReadSlice效率相近，要明显快于ReadString和ReadBytes。
*/

func ReadLine1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}

func ReadLine2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}

func ReadLine3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, _, err := reader.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
	}
}

func ReadLine4(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadSlice('\n')
		if err != nil || err == io.EOF {
			break
		}
	}
}

func ScanFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = scanner.Text()
	}
}

func main() {

}
