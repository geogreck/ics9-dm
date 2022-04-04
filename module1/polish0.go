package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// <Expr> ::= '(' <Sign> <Expr> <Expr> ')' | <Number>
// <Sign> ::= '+' | '-' | '*'
// <Number> ::= '0' | ... | '9'

func parseInt(expr []rune, i int) (int, int) {
	buf := make([]rune, 0)
	for expr[i] != ' ' && expr[i] != ')' {
		buf = append(buf, expr[i])
		i++
	}
	num, _ := strconv.Atoi(string(buf))
	return num, i
}

func count_rec(expr []rune, ind int) (int, int) {
	if ind >= len(expr) {
		return 0, ind
	}
	fmt.Println(string(expr[ind:]))
	if expr[ind] == '(' {
		if expr[ind+1] == '*' {
			i := ind + 3
			mul := 1
			for i < len(expr) && expr[i] != ')' {
				if unicode.IsDigit(expr[i]) {
					var a int
					a, i = parseInt(expr, i)
					mul *= a
					i++
				} else if expr[i] == '(' {
					var buf int
					buf, i = count_rec(expr, i)
					mul *= buf
				} else {
					i++
				}
			}
			return mul, i
		} else if expr[ind+1] == '+' {
			i := ind + 3
			sum := 0
			for i < len(expr) && expr[i] != ')' {
				if unicode.IsDigit(expr[i]) {
					var a int
					a, i = parseInt(expr, i)
					sum += a
					i++
				} else if expr[i] == '(' {
					var buf int
					buf, i = count_rec(expr, i)
					sum += buf
				} else {
					i++
				}
			}
			return sum, i
		} else {
			fmt.Println("I'm here")
			return 0, ind
		}
	} else {
		fmt.Println("Error")
	}
	return 0, 100
}

func count(expr string) int {
	ans, _ := count_rec([]rune(expr), 0)
	return ans
}

func main() {
	/* scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text() */
	exprs := []string{"", "(+ 1 2)", ")", "(* 3 4)", "(* 10 10 10)", "(* 5 (+ 3 4))", "(* (* 10 1) 1 (+ 10 15) 10)"}
	for _, expr := range exprs {
		ans := count(expr)
		fmt.Printf("%s = %d\n\n\n", expr, ans)
	}
}
