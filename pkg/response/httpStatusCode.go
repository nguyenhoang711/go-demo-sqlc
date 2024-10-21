package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeParamInvalid = 20003 //Email invalid

	ErrCodeNotAuthen = 20004 //Not authen
	ErrIOTPInvalid = 30002 // create otp failed
	ErrSendEmailOTPFail = 30003 // fail to send OTP
	
	// user
	ErrCodeUserEmailExist = 20005 // user email exits

	//conflict
	ErrCodeConflict = 40005 // internal server error
)

// message
func ErrMessageDict() func(int) string {
	msgs := map[int]string{
		ErrCodeSuccess: "success",
		ErrCodeParamInvalid: "email is invalid",
		ErrCodeNotAuthen: "you don't have access to this API",
		ErrIOTPInvalid: "OTP IS NOT VALID",
		ErrSendEmailOTPFail: "Send email otp failed",
		ErrCodeUserEmailExist: "user has already registered",
		ErrCodeConflict: "Internal server error",
	}

	return func(key int) string {
		return msgs[key]
	}
}