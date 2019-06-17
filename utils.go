package utils

import "time"

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

/*GetCurrentDateTime returns the current date/time
 * Returns: string - Current date/time in YYYY-mm-dd HH:ii::ss format
 */
func GetCurrentDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

// IsEmpty checks if a required string field is empty
func IsEmpty(field string) bool {
	return field == ""
}