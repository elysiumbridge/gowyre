package gowyre

type WyreError struct {
	Language    string `json:"language"`
	ExceptionId string `json:"exceptionId"`
	ErrorCode   string `json:"errorCode"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Transient   bool   `json:"transient"`
}

func (e WyreError) Error() string {
	return e.Message
}

type Error struct {
	WyreError
	StatusCode int    `json:"code"`
	Details    string `json:"details"`
}

func (e Error) Error() string {
	return e.Message
}

func NewValidationException(err WyreError) Error {
	return Error{err, 400, "The action failed due to problems with the request."}
}

func NewInsufficientFundsException(err WyreError) Error {
	return Error{err, 400, "You requested the use of more funds in the specified currency than were available"}
}

func NewAccessDeniedException(err WyreError) Error {
	return Error{err, 401, "You lack sufficient privilege to perform the requested action"}
}

func NewTransferException(err WyreError) Error {
	return Error{err, 400, "There was a problem completing your transfer request"}
}

func NewMFARequiredException(err WyreError) Error {
	return Error{err, 400, "An MFA action is required to complete the request. In general you should not get this exception while using API keys"}
}

func NewCustomerSupportException(err WyreError) Error {
	return Error{err, 400, "Please contact us at support@sendwyre.com to resolve this!"}
}

func NewNotFoundException(err WyreError) Error {
	return Error{err, 404, "You referenced something that could not be located"}
}

func NewRateLimitException(err WyreError) Error {
	return Error{err, 429, "Your requests have exceeded your usage restrictions. Please contact us if you need this increased"}
}

func NewAccountLockedException(err WyreError) Error {
	return Error{err, 400, "The account has had a locked placed on it for potential fraud reasons."}
}

func NewLockoutException(err WyreError) Error {
	return Error{err, 403, "The account or IP has been blocked due to detected malicious behavior"}
}

func NewUnknownException(err WyreError) Error {
	return Error{err, 500, "A problem with our services internally. This should rarely happen"}
}

func NewApiException(err WyreError) Error {
	return Error{err, 500, "A problem with our services internally. This should rarely happen"}
}

func NewError(err WyreError) Error {
	return map[string]Error{
		"ValidationException":        NewValidationException(err),
		"InsufficientFundsException": NewInsufficientFundsException(err),
		"AccessDeniedException":      NewAccessDeniedException(err),
		"TransferException":          NewTransferException(err),
		"MFARequiredException":       NewMFARequiredException(err),
		"CustomerSupportException":   NewCustomerSupportException(err),
		"NotFoundException":          NewNotFoundException(err),
		"RateLimitException":         NewRateLimitException(err),
		"AccountLockedException":     NewAccountLockedException(err),
		"LockoutException":           NewLockoutException(err),
		"UnknownException":           NewUnknownException(err),
		"ApiException":               NewApiException(err),
	}[err.Type]
}
