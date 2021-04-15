package main

import "fmt"

//正常声明函数的参数形式，若多个参数类型相同，可仅在最后声明
func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

//go语言也支持多返回值，例如同时返回一个函数的结果和错误信息
func vals() (int, int) {
	return 4, 13
}

//go语言也支持可变参数函数，在调用时可以传递任意数量的参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//go语言也支持递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	res := plus(1, 2)
	fmt.Println(res)
	res = plusplus(1, 2, 3)
	fmt.Println(res)
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	sum(1, 2)
	sum(1, 2, 3)
	//sum也可调用多个值的切片
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	fmt.Println(fact(7))
}
