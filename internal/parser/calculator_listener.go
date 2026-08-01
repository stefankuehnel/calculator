// Code generated from Calculator.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Calculator
import "github.com/antlr4-go/antlr/v4"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterCalculation is called when entering the calculation production.
	EnterCalculation(c *CalculationContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterUnary is called when entering the Unary production.
	EnterUnary(c *UnaryContext)

	// ExitCalculation is called when exiting the calculation production.
	ExitCalculation(c *CalculationContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitUnary is called when exiting the Unary production.
	ExitUnary(c *UnaryContext)
}
