package utils

import "github.com/go-playground/validator/v10"

func ErrorMessage(fe validator.FieldError) string {
	field := fe.Field()
	tag := fe.Tag()

	// --- Custom messages per field + tag (Register/Login/User) ---
	switch field {
	case "Email":
		if tag == "email" {
			return "invalid email format"
		}
	case "Password":
		if tag == "password_complex" {
			return field + " must contain uppercase, lowercase, number, and special character"
		}
	case "Username":
		if tag == "max" {
			return field + " must be at most " + fe.Param() + " characters"
		}
	}
	// --- Default messages ---
	switch tag {
	case "required":
		return field + " is required"
	case "gte":
		return field + " must be greater than or equal to " + fe.Param()
	case "max":
		return field + " can have at most " + fe.Param() + " item(s)"
	default:
		return field + " is invalid"
	}
}
