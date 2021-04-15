package main

import "fmt"

func main() {
	//创建切片需要使用内建函数make,例如新建一个长度为3，字符串类型的切片
	s := make([]string, 3)
	fmt.Println(s)
	//切片和数组一样赋初值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))

	//切片支持比数组更丰富的功能。支持内建函数append，例如：
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	//切片也支持复制,和切片
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	l := c[2:5]
	fmt.Println(l)
	l = c[:5]
	fmt.Println(l)
	l = c[2:]
	fmt.Println(l)

	//也可以直接给切片初始化赋初值
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	//切片也有多维结构
	twod := make([][]int, 3)
	for i := 0; i < 3; i++ {
		len1 := i + 1
		twod[i] = make([]int, len1)
		for j := 0; j < len1; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)
}
