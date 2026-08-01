package calculator

import (
	"strings"
	"testing"
)

// TestEvaluate tests the Evaluate function.
// The tests compare the results with the != operator and not with a
// tolerance. Each expected value in these tests is exact in a float64.
func TestEvaluate(t *testing.T) {
	t.Run("with an addition", func(t *testing.T) {
		// Arrange
		expression := "1 + 2"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 3.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with a subtraction", func(t *testing.T) {
		// Arrange
		expression := "10 - 4"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 6.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with operator precedence", func(t *testing.T) {
		// Arrange
		// A multiplication comes before an addition. Thus the result is
		// 7 and not 9.
		expression := "1 + 2 * 3"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 7.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with parentheses", func(t *testing.T) {
		// Arrange
		// The parentheses change the precedence. Thus the result is 9
		// and not 7.
		expression := "(1 + 2) * 3"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 9.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with a unary minus", func(t *testing.T) {
		// Arrange
		// A unary minus binds more than a multiplication. Thus the
		// expression is (-3) * 4.
		expression := "-3 * 4"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := -12.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with a division", func(t *testing.T) {
		// Arrange
		// The calculation uses a float64. Thus 7 / 2 gives 3.5 and not 3.
		expression := "7 / 2"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 3.5
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with a decimal number", func(t *testing.T) {
		// Arrange
		expression := "1.5 + 1.5"

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 3.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with whitespace", func(t *testing.T) {
		// Arrange
		// The lexer removes each space, tab and line break.
		expression := "\t1  +\n2 "

		// Act
		result, err := Evaluate(expression)

		// Assert
		if err != nil {
			t.Errorf("got error %q", err)
		}

		expectedResult := 3.0
		if result != expectedResult {
			t.Errorf("expected: %v; got: %v", expectedResult, result)
		}
	})

	t.Run("with a division by zero", func(t *testing.T) {
		// Arrange
		// A float64 gives +Inf for this expression. The user needs an
		// error and not +Inf.
		expression := "1 / 0"

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with an invalid character", func(t *testing.T) {
		// Arrange
		expression := "1 $ 2"

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with an incomplete expression", func(t *testing.T) {
		// Arrange
		expression := "1 +"

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with an unbalanced parenthesis", func(t *testing.T) {
		// Arrange
		expression := "(1 + 2"

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with a number that is too large", func(t *testing.T) {
		// Arrange
		// The lexer accepts this number, but it is too large for a
		// float64.
		expression := strings.Repeat("9", 400)

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("with an empty expression", func(t *testing.T) {
		// Arrange
		expression := ""

		// Act
		_, err := Evaluate(expression)

		// Assert
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}

// TestCalculatorListenerPop tests the pop method with an empty stack.
// A parse tree that is correct always puts a value on the stack before
// an exit method reads the value. Thus only a direct call starts this
// branch and holds the test coverage at 100 %.
func TestCalculatorListenerPop(t *testing.T) {
	t.Run("with an empty stack", func(t *testing.T) {
		// Arrange
		listener := &calculatorListener{}

		// Act
		value := listener.pop()

		// Assert
		if value != 0 {
			t.Errorf("expected: %v; got: %v", 0.0, value)
		}

		if listener.err == nil {
			t.Error("expected an error, got nil")
		}
	})
}
