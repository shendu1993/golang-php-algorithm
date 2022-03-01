package test

import (
	"fmt"
	"strconv"
)

//打印数据
func Dd(value interface{}) {
	fmt.Println("------------start------------")
	fmt.Println(value)
	fmt.Println("------------end--------------")
	fmt.Println("------------分割线------------")
}

//打印 循环次数
func DdCycleTimes(num int) {
	fmt.Println("Cycle " + strconv.Itoa(num) + " Times")
}
