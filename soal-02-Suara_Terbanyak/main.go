package main

import (
	"fmt"
	"slices"
)

func mayoritas(suara []int) int {
	if len(suara) == 1 {
		return suara[0]
	}
	threshold := float64(len(suara) / 2)
	fmt.Println("threshold", threshold)
	slices.Sort(suara)
	major := -1
	counter := 1
	current := suara[0]
	for i := 1; i < len(suara); i++ {
		fmt.Println("curr| counter",current,counter)
		if suara[i] == current {
			counter++
			if float64(counter) > threshold{
				major = current
			}
		} else {
			counter = 1
			current = suara[i]
		}
	}

	fmt.Println(suara)

	return major
}

func main() {
	suara := []int{1, 2, 3, 3, 3, 3}
	major:= mayoritas(suara)
	fmt.Println("major",major)
}
