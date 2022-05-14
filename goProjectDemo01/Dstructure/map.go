package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type MySet struct {
	nums map[int]bool
}

func NewMySet() *MySet {
	return &MySet{
		nums: make(map[int]bool),
	}
}

func (set *MySet) Add(num int) bool {
	if set == nil {
		return false
	}
	set.nums[num] = true
	return true
}

func (set *MySet) Remove(num int) bool {
	if set == nil {
		return false
	}
	delete(set.nums, num)
	return true
}

func (set *MySet) Size() int {
	if set == nil {
		return -1
	}
	return len(set.nums)
}

func (set *MySet) IsEmpty() bool {
	if set == nil {
		return true
	}
	return len(set.nums) == 0
}

func mapEquals(map1 map[string]int, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k1, v1 := range map1 {
		v2, ok := map2[k1]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func addEdge(from string, to string, graph map[string]map[string]bool) {
	if len(from) == 0 || len(to) == 0 {
		return
	}
	tos := graph[from]
	if tos == nil {
		tos = make(map[string]bool)
		graph[from] = tos
	}
	tos[to] = true
}

func checkEdgeExist(from string, to string, graph map[string]map[string]bool) bool {
	return graph[from][to]
}

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	number_counts := make(map[rune]int)
	letter_counts := make(map[rune]int)
	other_counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letter_counts[r]++
		} else if unicode.IsNumber(r) {
			number_counts[r]++
		} else {
			other_counts[r]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
