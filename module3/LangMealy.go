package main

import (
	"fmt"
	"sort"
)

type Automat struct {
	n, m, q          int
	transitionMatrix [][]int
	outputMatrix     [][]string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type StringSlice []string

func (s StringSlice) Len() int { return len(s) }

func (s StringSlice) Less(i, j int) bool {
	len1 := len(s[i])
	len2 := len(s[j])
	for k := 0; k < min(len1, len2); k++ {
		if []rune(s[i])[k] < []rune(s[j])[k] {
			return true
		} else if []rune(s[i])[k] > []rune(s[j])[k] {
			return false
		}
	}
	return false
}

func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (this *Automat) input() {
	var n, q int
	fmt.Scan(&n)
	this.n = n
	this.m = 2
	transitionMatrix := make([][]int, n)
	outputMatrix := make([][]string, n)
	for i := 0; i < n; i++ {
		transitionMatrix[i] = make([]int, 2)
		outputMatrix[i] = make([]string, 2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&transitionMatrix[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&outputMatrix[i][j])
		}
	}
	fmt.Scan(&q)
	this.q = q
	this.transitionMatrix = transitionMatrix
	this.outputMatrix = outputMatrix
}

func (this *Automat) display() {
	fmt.Printf("%d\n%d\n%d\n", this.n, this.m, this.q)
	for _, row := range this.transitionMatrix {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
	for _, row := range this.outputMatrix {
		for _, elem := range row {
			fmt.Printf("%v ", elem)
		}
		fmt.Println()
	}
}

func (this Automat) gen_words(rwords []string, length int, q int, depth_credit int) []string {
	changed := 0
	words1 := make([]string, 0)
	words2 := make([]string, 0)
	//fmt.Println("in: ", rwords)
	for j := range rwords {
		if len(rwords[j]) < length {
			k := rwords[j]
			temp := rwords[j]
			s := this.outputMatrix[q][0]
			if s != "-" {
				temp = temp + this.outputMatrix[q][0]
				changed++
			}
			temp2 := rwords[j]
			s = this.outputMatrix[q][1]
			if s != "-" {
				temp2 = temp2 + this.outputMatrix[q][1]
				changed++
			}
			//fmt.Println(this.outputMatrix[q][0], " | ", this.outputMatrix[q][1])
			//fmt.Println("q = ", q, "word1 = ", temp, " word2 = ", temp2)
			if temp != k || depth_credit > 0 {
				if temp == k {
					depth_credit--
				}
				words1 = append(words1, temp)
			}
			if temp2 != k || depth_credit > 0 {
				if temp2 == k {
					depth_credit--
				}
				words2 = append(words2, temp2)
			}
		}
	}
	//fmt.Printf("length = %d, changed = %d, words len = %d\n", length, changed, len(words))
	if changed > 0 {
		if len(words1) > 0 && len(words1[0]) < length {
			words1 = this.gen_words(words1, length, this.transitionMatrix[q][0], depth_credit)
		}
		if len(words2) > 0 && len(words2[0]) < length {
			words2 = this.gen_words(words2, length, this.transitionMatrix[q][1], depth_credit)
		}

		words1 = append(words1, words2...)
	}
	return words1
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	var automat Automat
	var m int
	automat.input()
	fmt.Scan(&m)
	var rwords []string
	for i := 1; i <= m; i++ {
		var words []string
		foo := ""
		words = append(words, foo)
		words = automat.gen_words(words, i, automat.q, automat.n*i)
		words = removeDuplicateStr(words)
		rwords = append(rwords, words...)
	}
	sort.Stable(StringSlice(rwords))
	for j := range rwords {
		fmt.Print(rwords[j], " ")
	}
	fmt.Println()
	//automat.display()
}
