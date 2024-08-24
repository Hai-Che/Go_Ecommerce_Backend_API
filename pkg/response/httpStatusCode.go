package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is not valid
	ErrCodeValidToken   = 30001 // Token is not valid
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",            // Success
	ErrCodeParamInvalid: "Email is not valid", // Email is not valid
	ErrCodeValidToken:   "Token is not valid", // Token is not valid
}
