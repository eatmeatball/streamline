grammar Ri;

ID: [a-zA-Z]+;
INT: [0-9]+;
NEWLINE:'\r'?'\n';
WS:[ \t]+->skip;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';



stat: expr NEWLINE # printExpr
    | ID '=' expr NEWLINE # assgin
    | NEWLINE # blank
    ;

expr: expr op=(MUL|DIV) expr # MulDiv
    | expr op=(ADD|SUB) expr # AddSub
    | INT # int
    | ID # id
    | '(' expr ')' # parens
    ;

prog: stat+ ;
