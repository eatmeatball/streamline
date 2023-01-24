grammar Ri;

// value

ECHO: 'echo';
ID: [a-zA-Z]+;
INT: [0-9]+;
NEWLINE:'\r'?'\n';
WS:[ \t]+->skip;

// op
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
MOD: '%';


prog: stat+ ;

stat: ECHO '(' ID ')' NEWLINE  # echoExpr
    | expr NEWLINE # printExpr
    | ID '=' expr NEWLINE # assgin
    | NEWLINE # blank
    ;


expr: expr MOD expr # Mod
    | expr op=(MUL|DIV) expr # MulDiv
    | expr op=(ADD|SUB) expr # AddSub
    | INT # int
    | ID # id
    | '(' expr ')' # parens
    ;

