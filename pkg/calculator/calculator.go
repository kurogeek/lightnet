package calculator

import "lightnet/pkg/entity"

type CalculatorService interface {
	Sum(input entity.Input) entity.Result
	Subtract(input entity.Input) entity.Result
	Multiply(input entity.Input) entity.Result
	Divide(input entity.Input) (entity.Result, error)
}
