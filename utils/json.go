package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "errors"
	

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	fmt.Println("Initializing utils package")
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	fmt.Println("WriteJsonResponse called")
	w.Header().Set("Content-Type", "application/json") // Set the content type to application/json

	w.WriteHeader(status) // Set the HTTP status code

	return json.NewEncoder(w).Encode(data) // Encode the data as JSON and write it to the response
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}

	response["status"] = "success"
	response["message"] = message
	response["data"] = data
	return WriteJsonResponse(w, status, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	fmt.Println("Error in WriteJsonErrorResponse:", err)
	response := map[string]any{}
	response["status"] = "error"
	response["message"] = message
	response["error"] = err.Error()
	fmt.Println("Response in WriteJsonErrorResponse:")
	return WriteJsonResponse(w, status, response)
}

func ReadJsonBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Prevent unknown fields from being included in the JSON body
	return decoder.Decode(result)
}

// func ReadJsonBody(r *http.Request, result any) error {
// 	decoder := json.NewDecoder(r.Body)
// 	decoder.DisallowUnknownFields()
// 	err := decoder.Decode(result)
// 	if err != nil {
// 		if err == io.EOF {
// 			return fmt.Errorf("request body is empty or malformed")
// 		}
// 		return fmt.Errorf("json decode error: %w", err)
// 	}
// 	return nil
// }

