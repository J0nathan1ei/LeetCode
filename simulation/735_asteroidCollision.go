package main

import "fmt"

// 模拟操作，主要是要理解题意
// 原理跟高中数学的分类讨论一样
// 时间:30min（在公司，有一定工作时长）
func asteroidCollision(asteroids []int) []int {
	for i := 0; i < len(asteroids)-1; i++ {
		if len(asteroids) <= 1 {
			break
		}
		if asteroids[i] > 0 && asteroids[i+1] < 0 {
			if asteroids[i] == -asteroids[i+1] {
				asteroids = append(asteroids[:i], asteroids[i+2:]...)
				// 当排除i所在的元素时，要判断前面有没有元素，来决定i的下一次位置
				if i-1 >= 0 {
					i -= 2
				} else {
					i -= 1
				}
			} else if asteroids[i] > -asteroids[i+1] {
				asteroids = append(asteroids[:i+1], asteroids[i+2:]...)
				i -= 1
			} else if asteroids[i] < -asteroids[i+1] {
				asteroids = append(asteroids[:i], asteroids[i+1:]...)
				if i-1 >= 0 {
					i -= 2
				} else {
					i -= 1
				}
			}
		}
	}
	return asteroids
}

func main() {
	d := []int{10, 2, -5}
	fmt.Println(asteroidCollision(d))
}
