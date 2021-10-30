package errcode

var (
	Success = NewError(0, "Success")
	ServerError = NewError(10000000, "Server internal error occurred")
	InvalidParams = NewError(10000001, "Parameter error")
	NotFound = NewError(10000002, "Specified file not found")
	UnauthorizedAuthNotExist = NewError(10000003, "Auth failed with provided AppKey and AppSecret")
	UnauthorizedTokenError = NewError(10000004, "Auth failed with invalid token")
	UnauthorizedTokenGenerate = NewError(10000005, "Auth failed, generate token failed")
	UnauthorizedTokenTimeout = NewError(10000007, "Auth failed because of timeout.")
	TooManyRequests = NewError(10000006, "Too many requests")
)
