package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

/*FindStringInSlice detects if a string is in a slice.

Receives:
	* str (string) - String that will be searched inside the slice
	* sl ([]string) - Slice of strings

Returns:
	* bool - Whether or not the string was found inside the slice
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

/*SetResponse sets the response to be sent to the user in any API endpoints.

Receives:
	* w (http.ResponseWriter) - Go's HTTP ResponseWriter struct
 	* statusCode (int) - HTTP status code of the response
 	* body (HTTPResponse) - Body of the HTTP response using a custom struct
*/
func SetResponse(w http.ResponseWriter, statusCode int, body HTTPResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
	return
}

// HandleError is a basic error handling method.
// Prints out the error that occurred and exits the application.
func HandleError(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}
