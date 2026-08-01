grammar Calculator;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
LPAREN: '(';
RPAREN: ')';
NUMBER: [0-9]+ ('.' [0-9]+)?;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
//
// The entry rule has the name "calculation" and not "start". The Go
// target of ANTLR changes a rule with the name "start" to the method
// Start_(), because the parser already has a method with the name
// Start(). The name "calculation" gives the clean method Calculation().
calculation: expression EOF;

// An alternative that comes first gets the higher precedence. Thus a
// multiplication comes before an addition, and "1 + 2 * 3" gives 7.
// The label after each alternative, for example "# MulDiv", makes the
// context type MulDivContext and the listener method ExitMulDiv.
expression:
	LPAREN expression RPAREN # Parens
	| SUB expression # Unary
	| expression op = (MUL | DIV) expression # MulDiv
	| expression op = (ADD | SUB) expression # AddSub
	| NUMBER # Number;
