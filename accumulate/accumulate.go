package accumulate

/*
func Accumulate(given []string, f func(string) string) []string {
	expected := make([]string, len(given))
	for i, s := range given {
		expected[i] = f(s)
	}
	return expected
}

type Converter func(string) string

func Accumulate(in []string, cn Converter) (out []string) {
	for _, s := range in {
		out = append(out, cn(s))
	}
	return
}
*/

// Accumulate applies a given operation function to every element of a given collection
func Accumulate(collection []string, operation func(string) string) []string {
	for i, item := range collection {
		collection[i] = operation(item)
	}
	return collection
}
