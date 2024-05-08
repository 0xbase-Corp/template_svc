package utils

// Unique takes a slice of strings and returns a new slice with duplicates removed.
func Unique(input []string) []string {
	seen := make(map[string]struct{})
	j := 0
	for _, val := range input {
		if _, ok := seen[val]; ok {
			continue
		}
		seen[val] = struct{}{}
		input[j] = val
		j++
	}

	return input[:j]
}
