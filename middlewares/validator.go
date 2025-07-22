package middlewares

import (
	// "AuthInGo/dto"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"context"
)

func RequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload any // Define the type of payload you expect

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body in requestvalidator", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)


		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}


func UserLoginValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO // Define the type of payload you expect

		// fmt.Println("UserLoginValidator middleware called")

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body in userloginvalidator", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		fmt.Println("Payload validated successfully:", payload)

		
		// Add the payload to the request context for further processing
		ctx := context.WithValue(r.Context(), "payload", payload)

		// fmt.Println("Context with payload set:", ctx)

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}


func UserCreateValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO // Define the type of payload you expect

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body in usercreatevalidator", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)


		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}