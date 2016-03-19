package main

import (
	"fmt"
	//"golang/peng/lesson"
)

func main() {
	/*
		array := [5]int{1, 2, 3, 4, 5}
		// 临时结构体数组
		structArray := []struct {
			name string
			age  int
		}{{"Tim", 18}, {"Jim", 20}}

		// 数组遍历
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
		for i, v := range structArray {
			fmt.Println(i, v)
		}
	*/

	// 4. 字符串
	var s1 string
	s1 = "abc"
	// 字符串连接
	s1 = s1 + "ddd中国"

	// 取长度
	n := len(s1)

	// 字节遍历
	for i := 0; i < n; i++ {
		ch := s1[i]
		fmt.Println(ch)
	}
	fmt.Println("--------------")
	// Unicode字符遍历
	for i, ch := range s1 {
		fmt.Println(i, ch, s1[i])
	}

	//lesson.TypeDemo()
}
