package main

import "fmt"

func maxArea(height []int) int {
	low, high := 0, len(height)-1
	var res int
	for low < high{
		num := min(height[low],height[high])
		res = max(res, num * (high - low))
		if height[low] < height[high]{
			low++
		} else{
			high--
		}
	}
	return res
}


func max(x, y int)int{
	if x > y{
		return x
	}
	return y
}

func min(x, y int)int{
	if x < y{
		return x
	}
	return y
}


func main(){
	d := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(d))
}