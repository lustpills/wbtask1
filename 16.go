package main

import (
	"fmt"
)

func partition(arr []int, l, r, c int) (int, int) {
	e := l
	g := r

	for i := l; i < g; {
		if arr[i] < c {
			arr[e], arr[i] = arr[i], arr[e]
			e++
			i++
		} else if arr[i] > c {
			g--
			arr[g], arr[i] = arr[i], arr[g]
		} else {
			i++
		}
	}

	return e, g
}

func quicksort(arr []int, l, r int) {
	if r-l >= 2 {
		c := arr[l+(r-l)/2]
		a, b := partition(arr, l, r, c)
		quicksort(arr, l, a)
		quicksort(arr, b, r)
	}
}

func main() {
	arr := []int{3, 5, 2, 8, 6, 4, 5, 7}
	l := 0
	r := len(arr)

	quicksort(arr, l, r)

	fmt.Println(arr)
}
