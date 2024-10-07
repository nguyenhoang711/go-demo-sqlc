package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeParamInvalid = 20003 //Email invalid

	ErrCodeNotAuthen = 20004 //Not authen
)

// message
func ErrMessageDict() func(int) string {
	msgs := map[int]string{
		ErrCodeSuccess: "success",
		ErrCodeParamInvalid: "email is invalid",
		ErrCodeNotAuthen: "you don't have access to this API",
	}

	return func(key int) string {
		return msgs[key]
	}
}