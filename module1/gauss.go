package main

import "fmt"

type Fraction struct {
	a, b int
}

type FractionMatrix struct {
	matrix [][]Fraction
	n      int
}

func (fracMat FractionMatrix) get(i, j int) Fraction {
	return fracMat.matrix[i][j]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return abs(a)
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func (f Fraction) Simplify() Fraction {
	a, b := f.a, f.b
	if a < 0 && b < 0 {
		a, b = -a, -b
	} else if a < 0 || b < 0 {
		a, b = -abs(a), abs(b)
	}
	del := gcd(a, b)
	return Fraction{a / del, b / del}
}

func add(f1, f2 Fraction) Fraction {
	denominator := lcm(f1.b, f2.b)
	numerator := f1.a*(denominator/f1.b) + f2.a*(denominator/f2.b)
	return Fraction{numerator, denominator}.Simplify()
}

func subtract(f1, f2 Fraction) Fraction {
	return add(f1, Fraction{-f2.a, f2.b})
}

func multiply(f1, f2 Fraction) Fraction {
	return Fraction{f1.a * f2.a, f1.b * f2.b}.Simplify()
}

func divide(f1, f2 Fraction) Fraction {
	return Fraction{f1.a * f2.b, f1.b * f2.a}.Simplify()
}

func inputMatrix(n int) FractionMatrix {
	var matrix [][]Fraction
	var a int
	for i := 0; i < n; i++ {
		matrix = append(matrix, []Fraction{})
		for j := 0; j < n+1; j++ {
			fmt.Scan(&a)
			matrix[i] = append(matrix[i], Fraction{a, 1})
		}
	}
	return FractionMatrix{matrix, n}
}

func (fracMat *FractionMatrix) Diagonalize() []int {
	track := make([]int, fracMat.n)
	for i := range track {
		track[i] = i
	}
	for i := 0; i < fracMat.n; i++ {
		leader := fracMat.matrix[i][track[i]]
		if leader.a == 0 {
			var newLeader int
			for j := i; j < fracMat.n; j++ {
				if fracMat.matrix[i][track[j]].a != 0 {
					newLeader = j
				}
			}
			if newLeader == 0 {
				return nil
			} else {
				leader = fracMat.matrix[i][track[newLeader]]
				track[i], track[newLeader] = track[newLeader], track[i]
			}
		}
		for j := 0; j < fracMat.n; j++ {
			fracMat.matrix[i][track[j]] = divide(fracMat.matrix[i][track[j]], leader)
		}
		fracMat.matrix[i][fracMat.n] = divide(fracMat.matrix[i][fracMat.n], leader)
		for j := i + 1; j < fracMat.n; j++ {
			leader = fracMat.matrix[j][track[i]]
			for k := 0; k < fracMat.n; k++ {
				fracMat.matrix[j][track[k]] = subtract(fracMat.matrix[j][track[k]],
					multiply(fracMat.matrix[i][track[k]], leader))
			}
			fracMat.matrix[j][fracMat.n] = subtract(fracMat.matrix[j][fracMat.n], multiply(fracMat.matrix[i][fracMat.n], leader))
		}
	}
	return track
}

func (fracMat FractionMatrix) Solve() []Fraction {
	track := fracMat.Diagonalize()
	if track == nil {
		return nil
	}

	solution := make([]Fraction, fracMat.n)
	for i := fracMat.n - 1; i >= 0; i-- {
		sum := fracMat.matrix[i][fracMat.n]
		for j := fracMat.n - 1; j > i; j-- {
			sum = subtract(sum, multiply(fracMat.matrix[i][track[j]], solution[track[j]]))
		}
		solution[track[i]] = divide(sum, fracMat.matrix[i][track[i]])
	}
	return solution
}

func (fracMat FractionMatrix) PrintMatrix() {
	for i := 0; i < fracMat.n; i++ {
		for j := 0; j < fracMat.n+1; j++ {
			fmt.Printf("%d/%d ", fracMat.get(i, j).a, fracMat.get(i, j).b)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)

	m := inputMatrix(n)
	solution := m.Solve()
	for _, x := range solution {
		fmt.Printf("%d/%d\n", x.a, x.b)
	}
}

/* 3
-4 -1 8 2
7 -7 7 3
5 -1 -4 7 */
/* 377/21
214/7
274/21 */
