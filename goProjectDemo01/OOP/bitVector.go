package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//add,remove ,exist ,string,elem  ,len,clear,Union,copy,intersection(交集)，subset（差集）,Symmetry difference(对称差)
const bits = 32 << (^uint(0) >> 63)

type IntSet struct {
	nums []uint
}

func NewIntSet() *IntSet {
	intset := new(IntSet)
	intset.nums = make([]uint, 0)
	return intset
}

func (set *IntSet) Add(val int) {
	word, bit := val/bits, val%bits
	for word > len(set.nums)-1 {
		set.nums = append(set.nums, 0)
	}
	set.nums[word] = set.nums[word] | (1 << uint(bit))
}

func (set *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		set.Add(num)
	}
}

func (set *IntSet) Remove(val int) {
	word, bit := val/bits, val%bits
	if word > len(set.nums)-1 {
		return
	}
	set.nums[word] = set.nums[word] & ^(1 << uint(bit))
}

func (set *IntSet) Exist(val int) bool {
	word, bit := val/bits, val%bits
	if word > len(set.nums)-1 {
		return false
	}
	u := (set.nums[word] & (1 << uint(bit)))
	if u == 0 {
		return false
	} else {
		return true
	}
}

//x 👉   word, bit (1<<bit)
//bits :   0 , 1 , 2  ...  62  63
//bit    63  62         i,     bit = 63 - j  word = i     x = i * bits + bit
func (set *IntSet) String() string {
	var buffer bytes.Buffer
	buffer.WriteByte('[')
	for i, num := range set.nums {
		for j := 0; j < bits; j++ {
			if num&(1<<j) != 0 {
				val := i*bits + j
				buffer.WriteString(strconv.Itoa(val))
				buffer.WriteByte(' ')
			}
		}
	}
	buffer.Truncate(buffer.Len() - 1)
	buffer.WriteByte(']')
	return buffer.String()
}

func (set *IntSet) Elems() []int {
	elems := make([]int, 0)
	for i, num := range set.nums {
		for j := 0; j < bits; j++ {
			if num&(1<<j) != 0 {
				elems = append(elems, i*bits+j)
			}
		}
	}
	return elems
}

func (set *IntSet) Len() int {
	len := 0
	for _, num := range set.nums {
		for j := 0; j < bits; j++ {
			if num&(1<<j) != 0 {
				len++
			}
		}
	}
	return len
}

func (set *IntSet) Clear() {
	set.nums = make([]uint, 0)
}

func (set *IntSet) Copy() *IntSet {
	nnums := make([]uint, len(set.nums))
	copy(nnums, set.nums)
	nset := new(IntSet)
	nset.nums = nnums
	return nset
}

func (set *IntSet) Union(tset *IntSet) {
	for len(set.nums) < len(tset.nums) {
		set.nums = append(set.nums, 0)
	}
	for i := range set.nums {
		set.nums[i] = set.nums[i] | tset.nums[i]
	}
	return
}

//A- B
//0 1    = 0
//0 0    = 0
//1 1    = 0
//1 0    = 1
func (set *IntSet) SubSet(tset *IntSet) {
	for i, _ := range set.nums {
		if i >= len(tset.nums) {
			break
		}
		for j := 0; j < bits; j++ {
			if tset.nums[i]&(1<<j) != 0 {
				set.nums[i] = set.nums[i] & (^(1 << j))
			}
		}
	}
	return
}

func (set *IntSet) InterSet(tset *IntSet) {
	for i, _ := range set.nums {
		set.nums[i] = set.nums[i] & tset.nums[i]
	}
	return
}

func (set *IntSet) SymSubSet(tset *IntSet) {

}

func main() {
	set1 := NewIntSet()
	/* 	set2 := NewIntSet() */
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)
	set1.AddAll(7, 8, 7, 1, 4, 8, 9, 2, 8)
	fmt.Printf("set1.String(): %v\n", set1.String())
}
