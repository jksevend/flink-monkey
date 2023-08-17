# FlinkMonkey language
Based on [Writing An Interpreter In Go](https://interpreterbook.com)

Run ``go run ./...``:
```
FlinkMonkey v1.0.0
>> var add = func(x,y) { x+y; };
{Type:VAR Literal:var}
{Type:IDENTIFIER Literal:add}
{Type:= Literal:=}
{Type:FUNCTION Literal:func}
{Type:( Literal:(}
{Type:IDENTIFIER Literal:x}
{Type:, Literal:,}
{Type:IDENTIFIER Literal:y}
{Type:) Literal:)}
{Type:{ Literal:{}
{Type:IDENTIFIER Literal:x}
{Type:+ Literal:+}
{Type:IDENTIFIER Literal:y}
{Type:; Literal:;}
{Type:} Literal:}}
{Type:; Literal:;}
```
