package response

const (
	ErrCodeSuccess      = 20001 //success
	ErrCodeParamInvalid = 20003 // email is invalid
)

// massage
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
}
