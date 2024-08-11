package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, numbersInSlices := range slices {
		if len(numbersInSlices) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbersInSlices[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
