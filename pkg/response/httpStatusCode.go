package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeParamInvalid = 2003 //Email invalid
)

// message
func ErrMessageDict() func(int) string {
	msgs := map[int]string{
		ErrCodeSuccess: "success",
		ErrCodeParamInvalid: "email is invalid",
	}

	return func(key int) string {
		return msgs[key]
	}
}