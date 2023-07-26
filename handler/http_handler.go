package handler

import (
	"log"
	"net/http"

	"github.com/giovanisilqueirasantos/coinconv/domain"
	"github.com/giovanisilqueirasantos/coinconv/service"
	"github.com/labstack/echo/v4"
)

type httpHandler struct {
	Service service.Service
}

func NewHttpHandler(e *echo.Echo, s service.Service) httpHandler {
	handler := httpHandler{Service: s}
	e.GET("/exchange/:amount/:from/:to/:rate", handler.HandleExchange)

	return handler
}

func (h httpHandler) HandleExchange(c echo.Context) error {
	amount := c.Param("amount")
	from := c.Param("from")
	to := c.Param("to")
	rate := c.Param("rate")

	var err error

	amountQuant := domain.Amount{}
	currencyFrom := domain.Currency{}
	currencyTo := domain.Currency{}
	rateAmount := domain.Amount{}

	err = amountQuant.New(amount)
	err = currencyFrom.New(from)
	err = currencyTo.New(to)
	err = rateAmount.New(rate)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	convAmount, err := h.Service.Exchange(c.Request().Context(), currencyFrom, currencyTo, amountQuant, rateAmount)
	if err != nil {
		log.Printf("Error trying to make exchange :%s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to do the exchange")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"valorConvertido": convAmount.Value, "simboloMoeda": currencyTo.Symbol})
}
