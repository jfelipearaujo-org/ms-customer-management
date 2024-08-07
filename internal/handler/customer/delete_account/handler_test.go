package delete_account_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jfelipearaujo-org/ms-customer-management/internal/handler/customer/delete_account"
	delete_account_svc "github.com/jfelipearaujo-org/ms-customer-management/internal/service/customer/delete_account"
	"github.com/jfelipearaujo-org/ms-customer-management/internal/shared/custom_error"
)

func TestHandler_Handle(t *testing.T) {
	t.Run("Should delete a customer", func(t *testing.T) {
		// Arrange
		service := delete_account_svc.NewMockService(t)

		service.On("Delete", mock.Anything, mock.Anything).
			Return(nil)

		handler := delete_account.NewHandler(service)

		reqBody := delete_account_svc.DeleteAccountRequest{
			Name:    "John Doe",
			Address: "Av. Brasil, 1000",
			Phone:   "1122334455",
		}

		body, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetPath("/api/v1/customers/delete-account")
		ctx.Set("userId", uuid.NewString())

		// Act
		err = handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.Code)
		service.AssertExpectations(t)
	})

	t.Run("Should return business error", func(t *testing.T) {
		// Arrange
		service := delete_account_svc.NewMockService(t)

		service.On("Delete", mock.Anything, mock.Anything).
			Return(custom_error.ErrRequestNotValid)

		handler := delete_account.NewHandler(service)

		reqBody := delete_account_svc.DeleteAccountRequest{
			Name:    "John Doe",
			Address: "Av. Brasil, 1000",
			Phone:   "1122334455",
		}

		body, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetPath("/api/v1/customers/delete-account")
		ctx.Set("userId", uuid.NewString())

		// Act
		err = handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusUnprocessableEntity, he.Code)
		assert.Equal(t, custom_error.AppError{
			Code:    http.StatusUnprocessableEntity,
			Message: "validation error",
			Details: "request not valid, please check the fields",
		}, he.Message)
		service.AssertExpectations(t)
	})

	t.Run("Should return internal error", func(t *testing.T) {
		// Arrange
		service := delete_account_svc.NewMockService(t)

		service.On("Delete", mock.Anything, mock.Anything).
			Return(assert.AnError)

		handler := delete_account.NewHandler(service)

		reqBody := delete_account_svc.DeleteAccountRequest{
			Name:    "John Doe",
			Address: "Av. Brasil, 1000",
			Phone:   "1122334455",
		}

		body, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetPath("/api/v1/customers/delete-account")
		ctx.Set("userId", uuid.NewString())

		// Act
		err = handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusInternalServerError, he.Code)
		assert.Equal(t, custom_error.AppError{
			Code:    http.StatusInternalServerError,
			Message: "internal error deleting customer",
			Details: "assert.AnError general error for testing",
		}, he.Message)
		service.AssertExpectations(t)
	})
}
