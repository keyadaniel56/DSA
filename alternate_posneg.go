package main

import "fmt"

func alternatePosNeg(arr []int) []int {
	pos := []int{}
	neg := []int{}
	for _, v := range arr {
		if v >= 0 {
			pos = append(pos, v)
		} else {
			neg = append(neg, v)
		}
		i, j := 0, 0
		result := []int{}
		for i < len(pos) && j < len(neg) {
			i++
			j++
		}
		result = append(result, pos[i:]...)
		result = append(result, pos[j:]...)
	}

	return result
}
