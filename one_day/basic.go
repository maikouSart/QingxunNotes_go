package main

import (
	"fmt"
)

// 结构体
type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

//基础语法
func basic() {
	fmt.Println("hello,world!")
	x := 3
	//1. 循环
	for x <= 3 {
		x += 1
		fmt.Print(x)
	}
	for y := 1; y < 3; y++ {
		fmt.Print(y)
	}
	//2. switch
	a := 2
	switch a {
	case 1:
		fmt.Print(a)
	default:
		fmt.Print(a)
	}
	//3.数组
	fmt.Print("数组（长度一定）")
	var b [5]int
	b[0] = 100
	fmt.Print(b)
	//4.切片
	s := make([]int, 3)
	s[0] = 1
	s = append(s, 2)
	fmt.Print(s)
	//5.map
	m := make(map[string]int)
	m["one"] = 1
	fmt.Print(m)
	//6.range
	nums := []int{1, 2, 3, 4}
	for i, j := range nums {
		fmt.Print(i, j)
	}
	v, ok := m["one"]
	fmt.Println(v, ok)
}

func add2pr(n *int) int {
	*n += 2
	return *n
}

func main() {
	// 1.基础语法
	//basic()
	// 2.指针
	n := 2
	add2pr(&n)
}
