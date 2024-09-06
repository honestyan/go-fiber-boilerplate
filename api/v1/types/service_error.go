package types

type ServiceError struct {
	Message string
	Code    int
	Error   error
}
