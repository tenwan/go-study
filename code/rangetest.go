package main

import "fmt"

func main() {
	//声明切片和数组的区别  数组nums:=[3]int{2,3,4}  切片nums:=[]int{2,3,4}
	nums := []int{2, 3, 4}
	sum := 0
	//range在数组和切片上提供对每项索引和值的访问，我们不需要索引，故使用空白标识符忽略
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	//下面为使用索引的情况
	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}
	//也可以使用range在map中迭代键值对
	kv := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kv {
		fmt.Println(k, v)
	}
	for k := range kv {
		fmt.Println(k)
	}
	for _, v := range kv {
		fmt.Println(v)
	}

}
