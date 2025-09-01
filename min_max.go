package main

import "fmt"

// return min and max
func min_max(array []int) (int, int) {
    min := array[0]
    max := array[0]

    for i := 1; i < len(array); i++ {
        if array[i] > max {
            max = array[i]
        }
        if array[i] < min {
            min = array[i]
        }
    }
    return min, max
}

func main() {
    array := []int{1, 7, 3, 4, 5}
    min, max := min_max(array)
    fmt.Println(min, max) // prints: 1 7
}
