package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
	"unicode"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

const (
	DefaultMaxLevel    int     = 18
	DefaultProbability float64 = 1 / math.E
)

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key   string
	value int
}

func (e *Element) Key() string {
	return e.key
}

func (e *Element) Value() int {
	return e.value
}

func (element *Element) Next() *Element {
	return element.next[0]
}

type SkipList struct {
	elementNode
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	prevNodesCache []*elementNode
}

func (list *SkipList) Front() *Element {
	return list.next[0]
}

func (list *SkipList) Set(key string, value int) *Element {
	var element *Element
	prevs := list.getPrevElementNodes(key)

	if element = prevs[0].next[0]; element != nil && element.key <= key {
		element.value = value
		return element
	}

	element = &Element{
		elementNode: elementNode{
			next: make([]*Element, list.randLevel()),
		},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = prevs[i].next[i]
		prevs[i].next[i] = element
	}

	list.Length++
	return element
}

func (list *SkipList) Get(key string) *Element {
	var prev *elementNode = &list.elementNode
	var next *Element

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && key > next.key {
			prev = &next.elementNode
			next = next.next[i]
		}
	}

	if next != nil && next.key <= key {
		return next
	}

	return nil
}

func (list *SkipList) Remove(key string) *Element {
	prevs := list.getPrevElementNodes(key)

	// found the element, remove it
	if element := prevs[0].next[0]; element != nil && element.key <= key {
		for k, v := range element.next {
			prevs[k].next[k] = v
		}

		list.Length--
		return element
	}

	return nil
}

func (list *SkipList) getPrevElementNodes(key string) []*elementNode {
	var prev *elementNode = &list.elementNode
	var next *Element

	prevs := list.prevNodesCache

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && key > next.key {
			prev = &next.elementNode
			next = next.next[i]
		}

		prevs[i] = prev
	}

	return prevs
}

func (list *SkipList) SetProbability(newProbability float64) {
	list.probability = newProbability
	list.probTable = probabilityTable(list.probability, list.maxLevel)
}

func (list *SkipList) randLevel() (level int) {
	r := float64(list.randSource.Int63()) / (1 << 63)

	level = 1
	for level < list.maxLevel && r < list.probTable[level] {
		level++
	}
	return
}

func probabilityTable(probability float64, MaxLevel int) (table []float64) {
	for i := 1; i <= MaxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

func NewWithMaxLevel(maxLevel int) *SkipList {
	if maxLevel < 1 || maxLevel > 64 {
		panic("maxLevel for a SkipList must be a positive integer <= 64")
	}

	return &SkipList{
		elementNode:    elementNode{next: make([]*Element, maxLevel)},
		prevNodesCache: make([]*elementNode, maxLevel),
		maxLevel:       maxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    DefaultProbability,
		probTable:      probabilityTable(DefaultProbability, maxLevel),
	}
}

func New() *SkipList {
	return NewWithMaxLevel(DefaultMaxLevel)
}

func lex(sentence string, array AssocArray) []int {
	expr := append([]rune(sentence), ' ')
	var buf []rune
	i := 1
	ind := make([]int, 0)
	for _, s := range expr {
		if unicode.IsSpace(s) {
			if len(buf) != 0 {
				x, exists := array.Lookup(string(buf))
				if !exists {
					array.Assign(string(buf), i)
					x = i
					i++
				}
				ind = append(ind, x)
				buf = []rune{}
			}
		} else {
			buf = append(buf, s)
		}
	}
	return ind
}

func (skl *SkipList) Lookup(s string) (x int, exists bool) {
	elem := skl.Get(s)
	if elem == nil {
		return 0, false
	}
	return elem.value, true
}

func (skl *SkipList) Assign(s string, x int) {
	skl.Set(s, x)
}

func main() {
	arr := AssocArray(New())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	ind := lex(s, arr)
	for _, x := range ind {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
