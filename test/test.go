package main

import (
	"fmt"
)

func main() {
	//指针取值
	// a := 10
	// b := &a // 取变量a的地址，将指针保存到b中
	// fmt.Printf("type of b:%T\n", b)
	// fmt.Printf("addr of b:%p\n", &b)
	// c := *b // 指针取值（根据指针去内存取值）
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)
	// fmt.Printf("addr of c:%p\n", &c)

	var num = 10
	fmt.Println("addr of num:%p", &num)
	fmt.Println("value of num:%p", num)
	var ptr *int
	ptr = &num
	*ptr = 11
	fmt.Println("addr of ptr:%p", ptr)
	fmt.Println("value of ptr:%p", *ptr)
	fmt.Println("value of num:%p", num)
}
