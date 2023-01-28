grammar Ri;

// Keywords
ECHO: 'echo';
IF:   'if';
BOOL: 'bool';
TRUE: 'true';
FALSE: 'false';
STRING: 'string';
FLOAT: 'float';
ANY: 'any';



ID: [a-zA-Z]+;
INT: [0-9]+;
NEWLINE:'\r'?'\n';
WS:[ \t]+->skip;
fragment DECIMALS
    : [0-9] ('_'? [0-9])*
    ;
DECIMAL_FLOAT_LIT : DECIMALS ('.' DECIMALS?);


// op
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
MOD: '%';


expr: expr MOD expr # Mod
    | expr op=(MUL|DIV) expr # MulDiv
    | expr op=(ADD|SUB) expr # AddSub
    | op=(FALSE|TRUE) # bool
    | INT # int
    | ID # id
    | '(' expr ')' # parens
    ;

parExpression
    : '(' expr ')'
    ;

stat: ECHO '(' ID ')' NEWLINE  # echoExpr
    | IF parExpression '{' NEWLINE? stat? '}' # ifLogic
    | expr NEWLINE # printExpr
    | ID '=' expr NEWLINE # assgin
    | NEWLINE # blank
    ;

prog: stat+ ;