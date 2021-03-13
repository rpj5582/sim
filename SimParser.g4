parser grammar SimParser;

options {
	tokenVocab = SimLexer;
}

start: (statement eos)*;

statement:
	LBRACE statement* RBRACE																# BlockStatement
	| IF expression statement																# IfStatement
	| LOOP statement																		# InfiniteLoopStatement
	| LOOP expression statement																# ConditionalLoopStatement
	| LOOP IDENTIFIER ASSIGNMENT min = expression TO max = expression statement				# LoopStatement
	| FUNCTION IDENTIFIER LPAREN (expression (COMMA expression)*)? RPAREN COLON IDENTIFIER	#
		FunctionStatement
	| type_ = IDENTIFIER varName = IDENTIFIER (
		ASSIGNMENT expression
	)?												# DeclarationStatement
	| varName = IDENTIFIER assignment_op expression	# AssignmentStatement
	| RETURN expression								# ReturnStatement
	| PRINT LPAREN expression RPAREN				# PrintStatement // TODO: remove this
	| RETURN										# ReturnStatement
	| BREAK											# BreakStatement
	| CONTINUE										# ContinueStatement;

expression:
	LPAREN expression RPAREN													# ParensExpression
	| SUBTRACT expression														# NegateExpression
	| NOT expression															# NotExpression
	| left = expression op = (MULTIPLY | DIVIDE | MODULO) right = expression	# MulDivModExpression
	| left = expression op = (ADD | SUBTRACT) right = expression				# AddSubExpression
	| left = expression op = (
		GREATER
		| LESSER
		| GREATER_OR_EQUAL
		| LESSER_OR_EQUAL
	) right = expression												# InequalityExpression
	| left = expression op = (EQUALS | NOT_EQUALS) right = expression	# EqualityExpression
	| left = expression AND right = expression							# AndExpression
	| left = expression OR right = expression							# OrExpression
	| IDENTIFIER														# VariableExpression
	| (NUMBER | TRUE | FALSE)											# LiteralExpression;

assignment_op:
	ASSIGNMENT
	| ADD_ASSIGNMENT
	| SUB_ASSIGNMENT
	| MUL_ASSIGNMENT
	| DIV_ASSIGNMENT
	| MOD_ASSIGNMENT;

eos:
	EOF
	| {lineTerminatorAhead(p)}?
	| {checkPreviousTokenText(p, "}")}?;