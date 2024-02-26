package domain

type ErrorCode string

const (
	InternalServerError ErrorCode = "UAMERR0001"
	BadRequestError     ErrorCode = "UAMERR0002"
)
