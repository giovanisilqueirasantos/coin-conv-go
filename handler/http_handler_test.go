package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/giovanisilqueirasantos/coinconv/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleExchangeWrongParams(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/exchange/:amount/:from/:to/:rate", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewHttpHandler(echo.New(), nil)

	handler.HandleExchange(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestHandleExchangeServiceError(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/exchange/:amount/:from/:to/:rate", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "real", "dollar", "4.50")

	mockService := new(service.MockService)
	mockService.On("Exchange", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error"))

	handler := NewHttpHandler(echo.New(), mockService)

	handler.HandleExchange(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestHandleExchangeSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/exchange/:amount/:from/:to/:rate", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "real", "dollar", "4.50")

	mockService := new(service.MockService)
	mockService.On("Exchange", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("45")

	handler := NewHttpHandler(echo.New(), mockService)

	handler.HandleExchange(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}
