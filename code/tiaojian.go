package main

import "fmt"

func main() {
	if 9%2 == 0 {
		fmt.Println("9是奇数")
	} else {
		fmt.Println("9是偶数")
	}
	//条件语句前可加声明语句
	if num := 9; num < 0 {
		fmt.Println(num, "1")
	} else if num < 10 {
		fmt.Println(num, "2")
	} else {
		fmt.Println(num, "3")
	}
}
