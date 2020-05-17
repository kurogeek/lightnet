package usecase

import (
	"lightnet/pkg/calculator"
	"lightnet/pkg/entity"
)

type calService struct {
}

func NewCalculatorService() calculator.CalculatorService {
	return &calService{}
}

func (cs calService) Sum(input entity.Input) entity.Result {
	var res entity.Result
	res.Result = input.A + input.B

	return res
}

func (cs calService) Subtract(input entity.Input) entity.Result {
	var res entity.Result
	res.Result = input.A - input.B

	return res
}

func (cs calService) Multiply(input entity.Input) entity.Result {
	var res entity.Result
	res.Result = input.A * input.B

	return res
}

func (cs calService) Divide(input entity.Input) (entity.Result, error) {
	var res entity.Result
	if input.B == 0 {
		return res, entity.ErrorCannotDivideByZero
	}
	res.Result = input.A / input.B

	return res, nil
}
