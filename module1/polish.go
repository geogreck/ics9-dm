package main

import (
	"fmt"
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
	symb := s.expr[s.i]
	return symb
}

// Продвижение вперёд
func (s *stream) next() rune {
	symb := s.expr[s.i]
	s.i++
	return symb
}

func count_rec(s stream) int {
	n := len(s.expr)
	if s.i >= n {
		return 0
	}
	if s.peek() == '*' {
		s.next()
		mul := 1
		for s.i < n && (unicode.IsDigit(s.peek()) || s.peek() == '*') {
			buf, _ := strconv.Atoi(string(s.next()))
			mul *= buf
		}
		if s.i < n && (s.peek() == '+' || s.peek() == '-') {
			return mul * count_rec(s)
		}
		return mul
	} else if s.peek() == '+' {
		s.next()
		sum := 0
		for s.i < n && (unicode.IsDigit(s.peek()) || s.peek() == '+') {
			buf, _ := strconv.Atoi(string(s.next()))
			sum += buf
		}
		if s.i < n && (s.peek() == '*' || s.peek() == '-') {
			return sum + count_rec(s)
		}
		return sum
	} else if s.peek() == '-' {
		s.next()
		sum := 0
		for s.i < n && (unicode.IsDigit(s.peek()) || s.peek() == '+') {
			buf, _ := strconv.Atoi(string(s.next()))
			sum -= buf
		}
		if s.i < n && (s.peek() == '*' || s.peek() == '+') {
			return sum - count_rec(s)
		}
		return sum
	} else {
		fmt.Printf("Error")
	}
	return 0
}

func count(expr string) int {
	return count_rec(stream{[]rune(expr), 0})
}

func test() bool {
	exprs := []string{"", "+12", "*34", "*999", "*5+34", "*91+78", "-10", "-56", "*90"}
	wants := []int{0, 3, 12, 729, 35, 135, -1, -11, 0}
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
	/* scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text() */
	fmt.Println(test())
}
