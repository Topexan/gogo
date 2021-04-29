package main

func sort(arr []int) []int {
	var c int
	for i := 0; i < len(arr) - 1; i++ {
		for j := i; j <len(arr); j++ {
			if arr[i] > arr[j] {
				c = arr[j]
				arr[j] = arr[i]
				arr[i] = c
			}
		}
	}
	return arr
}
