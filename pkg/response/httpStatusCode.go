package response

const (
	ErrCodeSuccess      = 20001 //success
	ErrCodeParamInvalid = 20003 // email is invalid
	ErrorInvalidToken   = 30001 // token invalid
)

// massage
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrorInvalidToken:   "token is invalid",
}
