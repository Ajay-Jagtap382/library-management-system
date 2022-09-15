package users

import "errors"

var (
	errEmptyID         = errors.New("transaction ID must be present")
	errNoTransaction   = errors.New("no Transaction present")
	errNoTransactionId = errors.New("transaction is not present")
)
