package helpers

type constantError string

func (e constantError) Error() string {
	return string(e)
}

var (
	ErrInternalServer  = constantError("internal server error")
	ErrIdIsNotFound    = constantError("id is not found")
	ErrBadRequest      = constantError("operation failed, request resource not valid")
	ErrForbiddenAccess = constantError("access denied!")
	ErrDataNotFound    = constantError("data not found")
	ErrBindData        = constantError("error bind data. data not valid")
	ErrReadingFile     = constantError("error reading image file")
)
