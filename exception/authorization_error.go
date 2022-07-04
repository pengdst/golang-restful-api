package exception

type AuthorizationError struct {
	Error string
}

func NewAuthorizationError(error string) AuthorizationError {
	return AuthorizationError{Error: error}
}
