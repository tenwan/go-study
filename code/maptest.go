package main

import "fmt"

func main() {
	//创建一个空map，可使用内建函数make(map[key-type]val-type)
	m := make(map[string]int)
	m["L1"] = 0
	m["L2"] = 1
	fmt.Println(m)
	x := m["L1"]
	fmt.Println(x)

	//内建函数len返回键值对数量
	fmt.Println(len(m))
	//内建函数delete删除键值对
	delete(m, "L2")
	fmt.Println(m)

	//从map取值，还有可以选择接收的第二个返回值，来判断map中是否存在该键，消除键不存在和键为0产生的歧义
	_, prs := m["L2"]
	fmt.Println(prs)

	//对map进行赋初值并初始化的操作
	n := map[string]int{"K1": 1, "K2": 2}
	fmt.Println(n)
}
