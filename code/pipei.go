package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("这是周末")
	default:
		fmt.Println("这是周内")
	}

	//可以使用不带表达式的switch实现if/else逻辑的另一种
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("这是上午")
	case t.Hour() > 12:
		fmt.Println("这是下午")
	}

	//类型开关可以比较类型，而非值，常用来发现接口值的类型
	whatAml := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("is bool")
		case int:
			fmt.Println("is int")
		default:
			fmt.Println("dont have this type")

		}
	}
	whatAml(true)
	whatAml(1)
	whatAml("sjdhd")

}
