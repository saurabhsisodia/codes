package main

import "fmt"

func main() {

	petrol := []int{4, 6, 7, 4}
	distance := []int{6, 5, 3, 5}

	prev, curr := 0, 0
	ans := 0

	for i := 0; i < len(petrol); i++ {
		curr = curr + petrol[i] - distance[i]
		if curr < 0 {
			prev = prev + curr
			curr = 0
			ans = i + 1
		}
	}
	if curr+prev >= 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}
