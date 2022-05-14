package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

//一.顶层接口语义： 对原来的slice进行改造，返回的还是原来的slice【尽量不要在底层新创建数组好吧】
//，需要将原来的slice进行更新  slice = append(slice, num)

func reverseSlice(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate1(nums []int, n int) []int {
	if len(nums) == 0 || n < 0 || n >= len(nums) {
		return nums
	}
	res := make([]int, len(nums))
	for i, _ := range nums {
		res[i] = nums[(i+n)%len(nums)]
	}
	return res
}

//将一个slice左移n个元素
func rotate2(nums []int, n int) []int {
	if len(nums) == 0 || n < 0 || n >= len(nums) {
		return nums
	}
	reverseSlice(nums[0:n])
	reverseSlice(nums[n:])
	reverseSlice(nums[:])
	return nums
}

func reverseArray(array *[9]int) {
	l := len(*array)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}

}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func removeSpace1(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	i := 0
	for _, str := range strs {
		if str != "" {
			strs[i] = str
			i++
		}
	}
	return strs[:i]
}

func removeSpace2(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	res := strs[:0]
	for _, str := range strs {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

//copy(nums[i:], nums[i+1:])
//0 1 2 3    i  i+1    n-1
func removeByIndex(nums []int, i int) []int {
	if len(nums) == 0 {
		return nums
	}
	if i < 0 || i >= len(nums) {
		return nums
	}
	copy(nums[i:], nums[i+1:])
	nums = nums[:len(nums)-1]
	return nums

}

func removeDuplicate(strs []string) []string {
	res := strs[:0]
	for i := 0; i < len(strs); {
		res = append(res, strs[i])
		j := i
		for ; j < len(strs) && strs[j] == strs[i]; j++ {
		}
		if j >= len(strs) {
			break
		}
		i = j
	}
	return res
}

func removeUnicodeSpace(bytes []byte) []byte {
	if len(bytes) == 0 {
		return bytes
	}
	index := 0
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		if !unicode.IsSpace(r) {
			copy(bytes[index:], bytes[i:i+size])
			index += size
			i = i + size
		} else {
			bytes[index] = ' '
			index++
			j := i
			r1, size1 := utf8.DecodeRune(bytes[j:])
			for unicode.IsSpace(r1) {
				j += size1
				r1, size1 = utf8.DecodeRune(bytes[j:])
			}
			if j >= len(bytes) {
				break
			}
			i = j
		}
	}
	return bytes[:index]
}

func reverseRuneInString(bytes []byte) []byte {
	if len(bytes) == 0 {
		return bytes
	}
	for i, j := 0, len(bytes)-1; i < j; {
		_, size_l := utf8.DecodeRune(bytes[i:])
		_, size_r := utf8.DecodeLastRune(bytes[:j+1])
		if size_l == size_r {
			swapBytes(bytes, i, i+size_l-1, j-size_r+1, j)
		} else if size_l < size_r {
			swapBytes(bytes, i, i+size_r-1, j-size_r+1, j)
			reverseBytes(bytes, j-size_r+1, j-size_r+size_l)
			reverseBytes(bytes, j-size_r+size_l+1, j)
			reverseBytes(bytes, j-size_r+1, j)
		} else if size_l > size_r {
			swapBytes(bytes, i, i+size_l-1, j-size_l+1, j)
			reverseBytes(bytes, i, i+size_r-1)
			reverseBytes(bytes, i+size_r, i+size_l-1)
			reverseBytes(bytes, i, i+size_l-1)
		}
		i += size_l
		j -= size_r
	}
	return bytes
}

func swapBytes(bytes []byte, l1 int, r1 int, l2 int, r2 int) {
	j := l2
	for i := l1; i <= r1; i, j = i+1, j+1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return
}

func reverseBytes(bytes []byte, left int, right int) {
	for i, j := left, right; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return
}
func main() {
	var r rune = '中'
	bs := []byte(string(r))
	fmt.Printf("%q\n", bs)
}
