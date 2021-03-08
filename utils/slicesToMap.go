package utils

import(
	"errors"
)

// Mapify takes in two slices and combines them into a single map
func Mapify(a []string, b []string) (map[string]string, error) {
	// get the lengths to test for length so you can
	// avoid pointless errors
	la := len(a)
	lb := len(b)
	if la != lb {
		return nil, errors.New("The length of the slices supplied don't match")
	}

	m := make(map[string]string)
	for i := range a {
		// the key and maps are set using
		// value of each slice
		m[a[i]] = b[i]

	}
	return m, nil
}