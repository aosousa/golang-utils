package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

/*
FindStringInSlice detects if a string is in a slice.

Receives:
  - str (string) - String that will be searched inside the slice
  - sl ([]string) - Slice of strings

Returns:
  - bool - Whether or not the string was found inside the slice
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

/*
SetResponse sets the response to be sent to the user in any API endpoints.

Receives:
  - w (http.ResponseWriter) - Go's HTTP ResponseWriter struct
  - statusCode (int) - HTTP status code of the response
  - body (HTTPResponse) - Body of the HTTP response using a custom struct
*/
func SetResponse(w http.ResponseWriter, statusCode int, body HTTPResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
	return
}

/*
CalculateStringDistance calculates the number of empty characters required to match a
longer string above it (i.e. in a printed table)

Receives:
  - baseDistance (int) - Length of the string to match
  - name (string) - String to add empty characters to

Returns:
  - string - New string with the required added spaces to match the string above it
*/
func CalculateStringDistance(baseDistance int, name string) string {
	stringDistanceLen := baseDistance - len(name)
	stringDistance := ""
	for i := 0; i <= stringDistanceLen; i++ {
		stringDistance = stringDistance + " "
	}

	return stringDistance
}

// HandleError is a basic error handling method.
// Prints out the error that occurred and exits the application.
func HandleError(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}

// LogMessage simply logs a message to the console with the current date/time.
func LogMessage(message string) {
	fmt.Printf("[%s] %s", GetCurrentDateTime(), message)
}
