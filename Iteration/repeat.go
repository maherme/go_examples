package iteration

// Repeat creates a string with a character repeated the number of times passed as repetitions.
func Repeat(character string, repetitions int) string {
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return repeated
}
