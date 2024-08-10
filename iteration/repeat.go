package iteration

func Repeat(a string, repeatNo int) (repeated string) {
	for i := 0; i < repeatNo; i++ {
		repeated += a
	}
	return
}
