grammar Ri;

// 保留关键字
//import, fragment, lexer, parser, grammar, returns,
//locals, throws, catch, finally, mode, options, tokens
// target 1
/**
return 1;
int i = 1;
i = 1+1;
int i2 = 2+2;
string s = "1"
float f = 1.1
bool b = true

Print(1);
Print(a);
Print("xxxxxx");
Print("xxx");
Print(a+1);
Print(a+a);
Print(s)

func functionName(){
    Print("this is a function")
}
func functionName2() int {
    return 1;
}
Print(functionName2());


*/

/**
func numberFunction() int {
    funcLogic
    return result
}
*/

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

FUNC: 'func';
INT: 'int';
FLOAT: 'float';
STRING: 'string';
BOOL:'bool';
ANY: 'any';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
RETURN: 'return';
IDENTIFIER: Letter LetterOrDigit*;

// 表达
// 表达 (*|/) 表达 乘除
// 表达 (+|-) 表达 加减
// number 数字
// 数字的优先级最低，所以会优先匹配出第一个 number (*|/) number 之后才会是别的
expression
   : expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | NUMBER                             # Number
   ;

// 下一步 加入参数， 支持自定义返回类型
// 多行运算
functionDeclaration
    : FUNC IDENTIFIER '(' ')' INT '{'
     RETURN expression
     '}'
    ;

blockStatement:
functionDeclaration
;

blockStatements:
 blockStatement* #BlockStms
;


start : blockStatements EOF;
;