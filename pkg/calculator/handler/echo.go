package handler

import (
	"lightnet/pkg/calculator"
	"lightnet/pkg/entity"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CalHandler struct {
	CalService calculator.CalculatorService
}

type ResposeMessage struct {
	Error *ErrorMessage `json:"error,omitempty"`
	Data  *DataMessage  `json:"data,omitempty"`
}

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type DataMessage struct {
	Item interface{} `json:"item"`
}

func NewCalHandler(e *echo.Echo, cs calculator.CalculatorService, sKey string) {
	handler := &CalHandler{
		CalService: cs,
	}

	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == sKey, nil
	}))

	e.POST("/calculator.sum", handler.Sum)
	e.POST("/calculator.sub", handler.Subtract)
	e.POST("/calculator.mul", handler.Multiply)
	e.POST("/calculator.div", handler.Divide)
}

func (ch CalHandler) Sum(c echo.Context) error {
	input := new(entity.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, ResposeMessage{
			Error: &ErrorMessage{
				Code:    400,
				Message: err.Error(),
			},
		})
	}

	result := ch.CalService.Sum(*input)
	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Item: result,
		},
	})
}

func (ch CalHandler) Subtract(c echo.Context) error {
	input := new(entity.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, ResposeMessage{
			Error: &ErrorMessage{
				Code:    400,
				Message: err.Error(),
			},
		})
	}

	result := ch.CalService.Subtract(*input)
	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Item: result,
		},
	})
}

func (ch CalHandler) Multiply(c echo.Context) error {
	input := new(entity.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, ResposeMessage{
			Error: &ErrorMessage{
				Code:    400,
				Message: err.Error(),
			},
		})
	}

	result := ch.CalService.Multiply(*input)
	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Item: result,
		},
	})
}

func (ch CalHandler) Divide(c echo.Context) error {
	input := new(entity.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(400, ResposeMessage{
			Error: &ErrorMessage{
				Code:    400,
				Message: err.Error(),
			},
		})
	}

	result, err := ch.CalService.Divide(*input)
	if err != nil {
		return c.JSON(400, ResposeMessage{
			Error: &ErrorMessage{
				Code:    400,
				Message: err.Error(),
			},
		})
	}
	return c.JSON(200, ResposeMessage{
		Data: &DataMessage{
			Item: result,
		},
	})
}
