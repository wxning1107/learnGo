package main

import "fmt"

// 校验Country是否实现StringAble接口
var _ StringAble = (*Country)(nil)

type StringAble interface {
	ToString() string
}

type Country struct {
	Name string
}

type City struct {
	Name string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p StringAble) {
	fmt.Println(p.ToString())
}

// 接口编程一个demo
// 这种编程模式在Go 的标准库有很多的示例，最著名的就是 io.Read 和 ioutil.ReadAll 的玩法，其中 io.Read 是一个接口，你需要实现他的一个 Read(p []byte) (n int, err error) 接口方法，只要满足这个规模，就可以被 ioutil.ReadAll这个方法所使用。
// 这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation”
func main() {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
