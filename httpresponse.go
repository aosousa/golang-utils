package utils

// HTTPResponse is the struct used to send back JSON responses in the API
type HTTPResponse struct {
	Success bool        `json:"success"` // Whether the request was successful or not
	Rows    int         `json:"rows"`    // Number of rows in the Data interface response
	Data    interface{} `json:"data"`    // Interface that has the relevant information to return. Can be a slice, object, string, etc.
	Error   string      `json:"error"`   // Error message (in case an error occurred)
}
