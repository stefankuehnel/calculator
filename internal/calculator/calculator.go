// Package calculator evaluates an arithmetic expression.
// The package uses a parser that ANTLR makes from the grammar in
// grammar/Calculator.g4. The task "codegen" makes the parser, and each
// build with Nix makes it again.
package calculator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/stefankuehnel/calculator/internal/parser"
)

// errorListener keeps the first syntax error of the lexer or the parser.
// ANTLR writes each syntax error to os.Stderr and then continues with a
// tree that is not complete. A CLI must not write to os.Stderr behind
// the command. Thus the package removes the listener of ANTLR and adds
// this listener, which keeps the error for the caller.
type errorListener struct {
	*antlr.DefaultErrorListener

	err error
}

// SyntaxError keeps the first error and ignores each error after it.
// One mistake in an expression often makes more than one error. The
// first error shows the position of the mistake.
func (l *errorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	if l.err == nil {
		l.err = fmt.Errorf("syntax error at line %d:%d: %s", line, column, msg)
	}
}

// calculatorListener calculates the result while it walks the tree.
// The listener keeps a stack. A number goes on the stack. An operator
// takes its operands from the stack and puts the result back.
type calculatorListener struct {
	*parser.BaseCalculatorListener

	stack []float64
	err   error
}

// push puts a value on the stack.
func (l *calculatorListener) push(value float64) {
	l.stack = append(l.stack, value)
}

// pop takes the top value from the stack.
// If the stack is empty, pop keeps an error and gives 0. A tree that
// the parser accepts always has a value on the stack. Thus pop gives an
// error and does not stop the program.
func (l *calculatorListener) pop() float64 {
	if len(l.stack) == 0 {
		l.fail(errors.New("expression is not complete"))

		return 0
	}

	value := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return value
}

// fail keeps the first error and ignores each error after it.
func (l *calculatorListener) fail(err error) {
	if l.err == nil {
		l.err = err
	}
}

// ExitNumber puts the number of the token on the stack.
// The lexer accepts each number that has the correct shape, but a
// number with many digits is too large for a float64.
func (l *calculatorListener) ExitNumber(c *parser.NumberContext) {
	text := c.GetText()

	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		l.fail(fmt.Errorf("number %q is out of range", text))

		return
	}

	l.push(value)
}

// ExitUnary puts the negative value of its operand on the stack.
func (l *calculatorListener) ExitUnary(_ *parser.UnaryContext) {
	l.push(-l.pop())
}

// ExitMulDiv multiplies or divides the two operands on the stack.
func (l *calculatorListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalculatorParserMUL:
		l.push(left * right)
	case parser.CalculatorParserDIV:
		// A float64 gives +Inf or NaN for a division by zero. The user
		// needs an error and not one of these values.
		if right == 0 {
			l.fail(errors.New("division by zero"))

			return
		}

		l.push(left / right)
	}
}

// ExitAddSub adds or subtracts the two operands on the stack.
func (l *calculatorListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalculatorParserADD:
		l.push(left + right)
	case parser.CalculatorParserSUB:
		l.push(left - right)
	}
}

// Evaluate calculates the result of an arithmetic expression.
// The expression obeys the grammar in grammar/Calculator.g4. It holds
// numbers, the operators +, -, * and /, a unary minus, and parentheses.
// The function gives an error for an expression that the parser does
// not accept and for a division by zero.
func Evaluate(expression string) (float64, error) {
	syntaxErrors := &errorListener{}

	lexer := parser.NewCalculatorLexer(antlr.NewInputStream(expression))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(syntaxErrors)

	calculatorParser := parser.NewCalculatorParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	calculatorParser.RemoveErrorListeners()
	calculatorParser.AddErrorListener(syntaxErrors)

	tree := calculatorParser.Calculation()

	// Look for a syntax error before the walk. The parser repairs an
	// expression that is not correct and gives a tree for it. A walk on
	// that tree gives a result that looks correct but is wrong.
	if syntaxErrors.err != nil {
		return 0, syntaxErrors.err
	}

	listener := &calculatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if listener.err != nil {
		return 0, listener.err
	}

	return listener.pop(), nil
}
