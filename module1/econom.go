package main

import "fmt"

// Наверное я просто очень люблю Scheme
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

type stack struct {
	arr []int
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
}

func (s *stack) pop() int {
	i := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return i
}

func econom_rec(s stream) int {
	calculated := make(map[string]bool)
	var stack stack
	n := len(s.expr)
	var count int
	for i := 0; i < n; i++ {
		if s.peek() == '(' {
			stack.push(i)
		} else if s.peek() == ')' {
			expr := string(s.expr[stack.pop() : i+1])
			_, flag := calculated[expr]
			if !flag {
				calculated[expr] = true
				count++
			}
		}
		s.next()
	}
	return count
}

func econom(expr string) int {
	return econom_rec(stream{[]rune(expr), 0})
}

func test() bool {
	exprs := []string{"x", "($xy)", "($(@ab)c)", "(#i($jk))", "(#($ab)($ab))", "(@(#ab)($ab))",
		"(#($a($b($cd)))(@($b($cd))($a($b($cd)))))", "(#($(#xy)($(#ab)(#ab)))(@z($(#ab)(#ab))))"}
	wants := []int{0, 1, 2, 2, 2, 3, 5, 6}
	errors := 0
	for i, expr := range exprs {
		ans := econom(expr)
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
	fmt.Println(test())
}
