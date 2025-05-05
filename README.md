# Bistro-Matic

Simple calculator written in Golang.
Takes the expression to evaluate from stdin.

## Examples

Compile

```sh
go build -o calc .
```

Run

```sh
$> echo '3+6' | ./calc 0123456789 '()+-*/%' 3 ; echo
9
$> echo '3v6' | ./calc 0123456789 '{}vwxyz' 3 ; echo
9
$> echo '----++-6(12)' | ./calc 0123456789 '()+-*/%' 10 ; echo
syntax error
$> echo '----++-6*12' | ./calc 0123456789 '()+-*/%' 11 | cat -e ; echo
-72
$> echo '-(12-(4*32))' | ./calc 0123456789 '()+-*/%' 12 | cat -e ; echo
116
$> echo '-(e@-(;*!@))' | ./calc '0A@!;ie& ]' '()+-*/%' 12 | cat -e ;echo
echo ee
$> echo '-(12*(13+15/5*(6/(12+14%(30%5+(10*25)-46)+16)-20)/43)*20)*(-(12-98*42)*(16+63-50/3))' | ./calc 0123456789 '()+-*/%' 84 | cat -e ; echo
-744629760
$> echo '-(&!-(;*!@))' | ./calc '~^@!;i &[]' "()+-*/%" 13 | cat -e ; echo
ii
```
