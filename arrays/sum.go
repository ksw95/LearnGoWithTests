package main

func Sum(array []int) (sum int) {
	for _, number := range array {
		sum += number
	}
	return
}
