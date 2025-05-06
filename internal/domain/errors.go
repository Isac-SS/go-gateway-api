package domain

import "errors"

var (
	ErrAccountNotFound   = errors.New("Account not found")
	ErrDuplicatedAPIKey  = errors.New("Api key already exists")
	ErrInvoiceNotFound   = errors.New("Invoice not found")
	ErrUnauthorizedAcess = errors.New("Unauthorized acess")
)
