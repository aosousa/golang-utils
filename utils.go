package utils

/*FindStringInSlice detects if a string is in a slice.
 * Receives:
 * str (string) - String that will be searched inside the slice
 * sl ([]string) - Slice of strings
 *
 * Returns: bool - Whether or not the string was found inside the slice
 */
func FindStringInSlice(str string, sl []string) bool {
	for _, el := range sl {
		if str == el {
			return true
		}
	}

	return false
}
