package main

import (
	"errors"
	"fmt"
)

func main() {
	//info := Info{}
	//var v Visitor = &info
	//v = LogVisitor{v}
	//v = NameVisitor{v}
	//v = OtherThingsVisitor{v}
	//v.Visit(func(info *Info, err error) error {
	//	fmt.Println("开始赋值")
	//	info.Name = "wxning"
	//	info.Namespace = "prd"
	//	info.OtherThings = "I'm programming."
	//	return nil
	//})

	visitors := NewDecoratedVisitor(new(Info), initInfo())
	for i := range registerVisitors {
		visitors = registerVisitors[i](visitors)
	}
	visitors.Visit(func(info *Info, err error) error {
		return err
	})
	//info := Info{}
	//var v Visitor = &info
	//v = NewDecoratedVisitor(v, func(info *Info, err error) error {
	//	fmt.Println("开始赋值")
	//	info.Name = "wxning"
	//	info.Namespace = "prd"
	//	info.OtherThings = "I'm programming."
	//	return nil
	//}, func(info *Info, err error) error {
	//	fmt.Println("v2")
	//	fmt.Println(info.Name)
	//	fmt.Println(info.Namespace)
	//	fmt.Println(info.OtherThings)
	//	return nil
	//})
	//v.Visit(func(info *Info, err error) error {
	//	fmt.Println("v")
	//	fmt.Println(info.Name)
	//	fmt.Println(info.Namespace)
	//	fmt.Println(info.OtherThings)
	//	return nil
	//})
}

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	fmt.Println("info visitor前")
	err := fn(info, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("info visitor后")
	return err
}

// Name visitor
type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor fn前")
		err = fn(info, err)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("NameVisitor fn 后")

		return errors.New("错误发生在name了")
	})
}

// Other visitor
type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor fn前")
		err = fn(info, err)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("OtherThingsVisitor fn后 ")
		//return errors.New("err occ")

		return errors.New("错误发生在other了")
	})
}

// Log visitor
type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor fn前")
		err = fn(info, err)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("LogVisitor fn后")

		return errors.New("错误发生在log了")
	})
}

var registerVisitors = []func(visitor Visitor) Visitor{newNameVisitor, newOtherVisitor, newLogVisitor}

func newNameVisitor(v Visitor) Visitor {
	return &NameVisitor{v}
}
func newOtherVisitor(v Visitor) Visitor {
	return &OtherThingsVisitor{v}
}
func newLogVisitor(v Visitor) Visitor {
	return &LogVisitor{v}
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return fn(info, nil)
	})
}

func initInfo() VisitorFunc {
	return func(info *Info, err error) error {
		fmt.Println("开始赋值")
		info.Name = "wxning"
		info.Namespace = "prd"
		info.OtherThings = "I'm programming."
		return nil
	}
}

/*
info visitor前
开始赋值
NameVisitor fn 前
OtherThingsVisitor fn 前
LogVisitor fn	前
===============================> logVisitor
===============================> OtherThings=I'm programming.
===========================> Name=wxning, NameSpace=prd
info visitor后

*/
