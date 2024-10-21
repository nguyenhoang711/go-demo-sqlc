package response

// CustomError struct to represent errors with a code and message
type CustomError struct {
    Code    int
    Message string
}

// Implement the Error method so that CustomError satisfies the error interface
func (e CustomError) Error() string {
    return e.Message
}

// Function to create a new custom error
func NewCustomError(code int, message string) CustomError {
    return CustomError{
        Code:    code,
        Message: message,
    }
}