lexer grammar SimLexer;

// Keywords
FUNCTION: 'function';
IF: 'if';
LOOP: 'loop';
TO: 'to';
RETURN: 'return';
BREAK: 'break';
CONTINUE: 'continue';

TRUE: 'true';
FALSE: 'false';

AND: 'and';
OR: 'or';
NOT: 'not';

// TODO: remove this
PRINT: 'print';

MULTIPLY: '*';
DIVIDE: '/';
ADD: '+';
SUBTRACT: '-';
MODULO: '%';

ASSIGNMENT: '=';
ADD_ASSIGNMENT: '+=';
SUB_ASSIGNMENT: '-=';
MUL_ASSIGNMENT: '*=';
DIV_ASSIGNMENT: '/=';
MOD_ASSIGNMENT: '%=';

EQUALS: '==';
NOT_EQUALS: '!=';
GREATER: '>';
LESSER: '<';
GREATER_OR_EQUAL: '>=';
LESSER_OR_EQUAL: '<=';

LPAREN: '(';
RPAREN: ')';

LBRACE: '{';
RBRACE: '}';

COLON: ':';

COMMA: ',';

fragment LETTER: [a-z|A-Z] | '_';
fragment DIGIT: [0-9];

NUMBER: DIGIT+ ([.] DIGIT+)?;

IDENTIFIER: LETTER (LETTER | DIGIT)*;

NEWLINE: [\r\n]+ -> channel(HIDDEN);
WHITESPACE: [ \t]+ -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);
BLOCK_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);