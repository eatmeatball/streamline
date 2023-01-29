grammar Ri;

// 赋值
// 循环
// 判断
// 函数
// Keywords
ECHO: 'echo';
IF:   'if';
FOR: 'for';
BOOL: 'bool';
STRING: 'string';
FLOAT: 'float';
ANY: 'any';
FUNC: 'func';
TRUE: 'true';
FALSE: 'false';


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

ifstat: stat+;

stat: ECHO '(' ID ')' NEWLINE  # echoExpr
    | IF parExpression '{' ifstat '}' # ifLogic
    | expr NEWLINE # printExpr
    | ID '=' expr NEWLINE # assgin
    | FOR '(' ';' ';' ')' '{'  stat+ '}' # forLogic
    | NEWLINE # blank
    ;

// for [( )]  [{  }]

prog: stat+ ;