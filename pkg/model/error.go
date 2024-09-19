package model

// Add this new struct for custom error messages
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Update the ErrorResponse struct
type ErrorResponse struct {
	Error   string            `json:"error" example:"Validation failed"`
	Details []ValidationError `json:"details,omitempty"`
}
