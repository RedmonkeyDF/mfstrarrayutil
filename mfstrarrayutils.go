package mfstrarrayutil

func StringArrayIndex(haystack []string, needle string) int{

	for i, v := range haystack {
		if v == needle {

			return i
		}
	}
	return -1
}

func StringArrayContains(haystack []string, needle string) bool {

	return StringArrayIndex(haystack, needle) > -1
}

func StringArrayMap(haystack []string, f func(string) string) []string {

	lh := len(haystack)
	if lh == 0 {

		return []string{}
	}

	outarr := make([]string, lh)

	for idx, val := range haystack {

		outarr[idx] = f(val)
	}

	return outarr
}

func StringArrayFilter(haystack []string, f func(string) bool) []string {

	outarr := []string{}

	for _, val := range haystack {

		if f(val) {

			outarr = append(outarr, val)
		}
	}

	return outarr
}

func StringArrayCount(haystack []string, f func(string) bool) int {

	count := 0

	for _, val := range haystack {

		if f(val) {

			count++
		}
	}

	return count
}

func StringArrayAny(haystack []string,  f func(string) bool) bool {

	return StringArrayCount(haystack, f) > 0
}

func StringArrayAll(haystack []string,  f func(string) bool) bool {

	return StringArrayCount(haystack, f) == len(haystack)
}