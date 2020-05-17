package usecase

import (
	"lightnet/pkg/entity"
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func isEqual(a, b float64) bool {
	acceptableAccuracy := 0.000001

	return math.Abs(a-b) <= acceptableAccuracy
}

func TestSum(t *testing.T) {
	cs := NewCalculatorService()

	t.Run("success-normal", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: 1,
		}
		expected := entity.Result{
			Result: 2,
		}
		res := cs.Sum(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-zero", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -1,
		}
		expected := entity.Result{
			Result: 0,
		}
		res := cs.Sum(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -20,
		}
		expected := entity.Result{
			Result: -19,
		}
		res := cs.Sum(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus-positive", func(t *testing.T) {
		input := entity.Input{
			A: 100,
			B: -99,
		}
		expected := entity.Result{
			Result: 1,
		}
		res := cs.Sum(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-with-real-number", func(t *testing.T) {
		input := entity.Input{
			A: 1.1,
			B: 2.2,
		}
		expected := entity.Result{
			Result: 3.3,
		}
		res := cs.Sum(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})
}

func TestSubtract(t *testing.T) {
	cs := NewCalculatorService()

	t.Run("success-normal", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: 1,
		}
		expected := entity.Result{
			Result: 0,
		}
		res := cs.Subtract(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-zero", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -1,
		}
		expected := entity.Result{
			Result: 2,
		}
		res := cs.Subtract(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -20,
		}
		expected := entity.Result{
			Result: 21,
		}
		res := cs.Subtract(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus-positive", func(t *testing.T) {
		input := entity.Input{
			A: 100,
			B: -99,
		}
		expected := entity.Result{
			Result: 199,
		}
		res := cs.Subtract(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-with-real-number", func(t *testing.T) {
		input := entity.Input{
			A: 1.1,
			B: 2.2,
		}
		expected := entity.Result{
			Result: -1.1,
		}
		res := cs.Subtract(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})
}

func TestMultiply(t *testing.T) {
	cs := NewCalculatorService()

	t.Run("success-normal", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: 1,
		}
		expected := entity.Result{
			Result: 1,
		}
		res := cs.Multiply(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-zero", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -1,
		}
		expected := entity.Result{
			Result: -1,
		}
		res := cs.Multiply(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -20,
		}
		expected := entity.Result{
			Result: -20,
		}
		res := cs.Multiply(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus-positive", func(t *testing.T) {
		input := entity.Input{
			A: 100,
			B: -99,
		}
		expected := entity.Result{
			Result: -9900,
		}
		res := cs.Multiply(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-with-real-number", func(t *testing.T) {
		input := entity.Input{
			A: 1.1,
			B: 2.2,
		}
		expected := entity.Result{
			Result: 2.42,
		}
		res := cs.Multiply(input)

		assert.Assert(t, isEqual(res.Result, expected.Result))
	})
}

func TestDivide(t *testing.T) {
	cs := NewCalculatorService()

	t.Run("success-normal", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: 1,
		}
		expected := entity.Result{
			Result: 1,
		}
		res, err := cs.Divide(input)

		assert.NilError(t, err)
		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-zero", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -1,
		}
		expected := entity.Result{
			Result: -1,
		}
		res, err := cs.Divide(input)

		assert.NilError(t, err)
		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus", func(t *testing.T) {
		input := entity.Input{
			A: 1,
			B: -20,
		}
		expected := entity.Result{
			Result: -0.05,
		}
		res, err := cs.Divide(input)

		assert.NilError(t, err)
		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-minus-positive", func(t *testing.T) {
		input := entity.Input{
			A: -100,
			B: -25,
		}
		expected := entity.Result{
			Result: 4,
		}
		res, err := cs.Divide(input)

		assert.NilError(t, err)
		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("success-with-real-number", func(t *testing.T) {
		input := entity.Input{
			A: 1.1,
			B: 2.2,
		}
		expected := entity.Result{
			Result: 0.5,
		}
		res, err := cs.Divide(input)

		assert.NilError(t, err)
		assert.Assert(t, isEqual(res.Result, expected.Result))
	})

	t.Run("fail-divide-by-zero", func(t *testing.T) {
		input := entity.Input{
			A: 10,
			B: 0,
		}

		res, err := cs.Divide(input)

		assert.Error(t, err, entity.ErrorCannotDivideByZero.Error())
		assert.DeepEqual(t, res, entity.Result{})
	})
}
