package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Прямо как в моем любимом Scheme
type stream struct {
	expr []rune
	i    int
}

// Запрос текущего символа
func (s stream) peek() rune {
	for s.i < len(s.expr) && (s.expr[s.i] == '(' || s.expr[s.i] == ')' || s.expr[s.i] == ' ') {
		s.i++
	}
	var symb rune
	if s.i == len(s.expr) {
		symb = 0
	} else {
		symb = s.expr[s.i]
	}
	return symb
}

// Продвижение вперёд
func (s *stream) next() rune {
	for s.i < len(s.expr) && (s.expr[s.i] == '(' || s.expr[s.i] == ')' || s.expr[s.i] == ' ') {
		s.i++
	}
	var symb rune
	if s.i == len(s.expr) {
		symb = 0
	} else {
		symb = s.expr[s.i]
	}
	s.i++
	return symb
}

func count_rec(s *stream) int {
	n := len(s.expr)
	s.peek()
	if s.i >= n {
		return 0
	}
	if s.peek() == '*' {
		s.next()
		mul := 1
		var count int
		for count = 0; count < 2; count++ {
			if unicode.IsDigit(s.peek()) {
				buf, _ := strconv.Atoi(string(s.next()))
				mul *= buf
			} else if s.peek() == '+' || s.peek() == '*' || s.peek() == '-' {
				mul *= count_rec(s)
			}
		}
		return mul
	} else if s.peek() == '+' {
		s.next()
		sum := 0
		var count int
		for count = 0; count < 2; count++ {
			if unicode.IsDigit(s.peek()) {
				buf, _ := strconv.Atoi(string(s.next()))
				sum += buf
			} else if s.peek() == '+' || s.peek() == '*' || s.peek() == '-' {
				sum += count_rec(s)
			}
		}
		return sum
	} else if s.peek() == '-' {
		s.next()
		sum := 0
		first := true
		var count int
		for count = 0; count < 2; count++ {
			if unicode.IsDigit(s.peek()) {
				buf, _ := strconv.Atoi(string(s.next()))
				if first {
					sum += buf
					first = false
				} else {
					sum -= buf
				}
			} else if s.peek() == '+' || s.peek() == '*' || s.peek() == '-' {
				if first {
					sum += count_rec(s)
					first = false
				} else {
					sum -= count_rec(s)
				}
			}
		}
		return sum
	} else {
		ans, _ := strconv.Atoi(string(s.peek()))
		return ans
	}
}

func count(expr string) int {
	str := stream{[]rune(expr), 0}
	return count_rec(&str)
}

func test() bool {
	exprs := []string{"", "+12", "*34", "*999", "*5+34", "*91+78", "-1", "-56", "*90", "-5*62", "+*34*23"}
	wants := []int{0, 3, 12, 729, 35, 135, -1, -1, 0, -7, 0}
	errors := 0
	for i, expr := range exprs {
		ans := count(expr)
		if ans == wants[i] {
			fmt.Printf("OK: %s = %d\n", expr, ans)
		} else {
			fmt.Printf("Error: %s = %d, Expected: %d\n", expr, ans, wants[i])
			errors++
		}
	}
	return errors == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(count(s))
}
