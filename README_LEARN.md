# Learn list

lexer : 词法分析程序
parser : 解析器

> 工欲善其事，必先利其器 

[Antlr](https://github.com/antlr/antlr4)
[Antlr官网](https://www.antlr.org/)
> ANTLR（语言识别的其他工具）是一个强大的「解析器」生成器，用于读取、处理、执行或翻译结构化文本或二进制文件。


```shell
antlr4 -Dlanguage=Go -o parser -visitor -listener  Ri.g4
```

[Antlr例子](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/)
[antlr/grammars-v4](https://github.com/antlr/grammars-v4)
