package errors

const (
	CodePositionNotFound = Code("ERR_POSITION_NOT_FOUND")
)

func NewPositionFoundError() *Error {
	return NewError("position not found", CodePositionNotFound)
}
