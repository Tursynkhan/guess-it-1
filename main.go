package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num int
	arr := []int{}
	lowRang := 0
	highRang := 0
	for i := 1; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)
		sort.Ints(arr)
		med := Median(arr)
		lowRang = int(med) - 45
		highRang = int(med) + 45
		fmt.Println(lowRang, highRang)
	}
}

func Median(data []int) float64 {
	l := len(data)
	if len(data)%2 == 0 {
		return float64(data[l/2]+data[l/2-1]) / 2
	} else {
		return float64(data[(l-1)/2])
	}
}
