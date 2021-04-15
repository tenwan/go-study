package main

import "fmt"

func main() {
	//声明大小为5的int类型数组
	var a [5]int
	fmt.Print(a)
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))

	//一维数组赋初值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//构造多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
