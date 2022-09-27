lexer grammar ApiLexer;

ATHANDLER:   '@handler';

WS:          [ \t\r\n\u000C]+ -> channel(88);
RAW_STRING:  '`' (~[`\\\r\n] | EscapeSequence)+ '`';
ID:          Letter LetterOrDigit*;

fragment EscapeSequence: '\\' [btnfr"'\\];
fragment LetterOrDigit: Letter | Digit;
fragment Digits: [0-9] ([0-9_]* [0-9])?;
fragment Digit: [0-9];
fragment Letter: [a-zA-Z_];