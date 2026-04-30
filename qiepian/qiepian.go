package main

import (
	"fmt"
)

// 第1种船舰切片的方式
func makeqiepian1() {
	//创建一个长度和容量都为5的int切片，元素初始化为零值
	s1 := make([]int, 5)

	//创建一个长度为3，容量为5的int切片
	s2 := make([]int, 3, 5)

	//打印切片
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s1长度", len(s1)) //长度
	fmt.Println("s1容量", cap(s1)) //容量
	fmt.Println("s2长度", len(s2))
	fmt.Println("s2容量", cap(s2))

	//格式化打印
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

// 第2种创建切片的方式
func makeqiepian2() {
	//长度和容量都为3
	s3 := []int{1, 2, 3}

	//通过索引指定位置初始化，长度为5，容量为5
	s4 := []int{0: 1, 4: 5}

	//打印
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
}

func main() {
	makeqiepian2()
}
