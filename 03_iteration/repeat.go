package iteration

// Given a string and repetition count copies the string and returns it.
// It panics if the repetitions is negative.
func Repeat(character string, repetitions int) string {
	if repetitions < 0 {
		panic("repetitions must be greater than 0")
	}
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return repeated
}
