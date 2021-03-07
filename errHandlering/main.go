package main

import (
	"encoding/binary"
	"io"
)

type Point struct {
	Longitude     interface{}
	Latitude      interface{}
	Distance      interface{}
	ElevationGain interface{}
	ElevationLoss interface{}
}

// 这种写法很丑
func parse(r io.Reader) (*Point, error) {

	var p Point

	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}

	return &p, nil
}

// 用函数式编程解决上述错误处理问题
// 我们通过使用Closure 的方式把相同的代码给抽出来重新定义一个函数，这样大量的  if err!=nil 处理的很干净了。但是会带来一个问题，那就是有一个 err 变量和一个内部的函数，感觉不是很干净。
func functionalErrHandleParse(r io.Reader) (*Point, error) {
	var p Point
	var err error

	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// 效仿bufio.Scanner()
/**
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
    	token := scanner.Text()
   	 // process token
	}
	if err := scanner.Err(); err != nil {
    	// process the error
	}
*/
type reader struct {
	reader io.Reader
	err    error
}

func NewReader(r io.Reader) *reader {
	return &reader{
		reader: r,
		err:    nil,
	}
}

func (r reader) Read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.reader, binary.BigEndian, data)
	}
}

func ScannerErrHandleParse(r io.Reader) (*Point, error) {
	var p Point
	reader := NewReader(r)

	reader.Read(&p.ElevationLoss)
	reader.Read(&p.ElevationGain)
	reader.Read(&p.Longitude)
	reader.Read(&p.Latitude)
	reader.Read(&p.Distance)

	if reader.err != nil {
		return nil, reader.err
	}

	return &p, nil
}

func main() {

}
