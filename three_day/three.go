package main

import (
	"fmt"
)

type Codec struct{}

//func Constructor() (_ Codec) { return }

// 实现结构体的方法
func (Codec) serialize(i int) string {
	fmt.Print(i)
	return "123"
}

type Animal interface {
	eat(i int)
}
type Man interface {
	eat(i int)
}
type Me struct {
	name int
	age  int
}

func (*Me) eat(i int) {
	fmt.Println(i)
}

func show(me *Me) {
	me.age = 100
}

func main() {
	//codec := Codec{}
	//codec.serialize(123)
	m := &Me{1, 2}
	show(m)
	fmt.Println(m.age)
}
