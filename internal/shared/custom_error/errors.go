package custom_error

import "net/http"

var (
	ErrRequestNotValid BusinessError = New(http.StatusUnprocessableEntity, "validation error", "request not valid, please check the fields")

	ErrCustomerNotFound BusinessError = New(http.StatusNotFound, "customer not found", "unable to find customer with the given id")

	ErrDeletionRequestAlreadyCreated BusinessError = New(http.StatusBadRequest, "deletion request already created", "deletion request already created for the given customer id")
	ErrDeletionRequestNotFound       BusinessError = New(http.StatusNotFound, "deletion request not found", "unable to find deletion request with the given customer id")
)
