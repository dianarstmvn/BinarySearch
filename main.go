package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Сколько чисел вы хотите найти?")
	var find int
	integ := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
		29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	fmt.Scan(&find)
	var target int
	fmt.Println("Введите числа через Enter, которые хотите найти")
	for i := 0; i < find; i++ {
		fmt.Scan(&target)
		go search(integ, target)
	}

	time.Sleep(10 * time.Millisecond)
}

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if target == nums[mid] {
			fmt.Println("Число ", target, " стоит на ", mid, " месте.")
			return 0
		}
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Числа ", target, "нет в массиве.")
	return 0
}
