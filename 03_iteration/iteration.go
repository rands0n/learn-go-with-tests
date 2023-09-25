package iteration

func Repeat(word string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += word
	}

	return repeated
}
