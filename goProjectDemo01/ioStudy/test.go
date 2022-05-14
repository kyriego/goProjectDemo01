package main

import (
	"bytes"
	"unicode"
)

//125215 ->125,215   str[:len-3]+","+str[len-3:]
func addcomma(str string) string {
	var buffer bytes.Buffer
	l := len(str)
	for i := 0; i < l; i++ {
		buffer.WriteByte(str[l-1-i])
		if (i+1)%3 == 0 && i != l-1 {
			buffer.WriteByte(',')
		}
	}
	l = buffer.Len()
	var bytes []byte = make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i], _ = buffer.ReadByte()
	}
	left, right := 0, l-1
	for left < right {
		swapBytes(bytes, left, right)
		left++
		right--
	}
	return string(bytes)
}

func swapBytes(bytes []byte, left int, right int) {
	bytes[left], bytes[right] = bytes[right], bytes[left]
}

func addcommaf(f string) string {
	if f == "" {
		return f
	}
	var n_flag bool
	var buffer bytes.Buffer
	if f[0] == '+' {
		f = f[1:]
	}
	if f[0] == '-' {
		n_flag = true
		f = f[1:]
	}
	l := len(f)
	count := 1
	for i := 0; i < l; i++ {
		buffer.WriteByte(f[l-1-i])
		if count%3 == 0 && i != l-1 {
			buffer.WriteByte(',')
		}
		if unicode.IsDigit(rune(f[l-1-i])) {
			count++
		}
	} //3.125125     buffer: 521521.3-    bytes: +
	if n_flag == true {
		buffer.WriteByte('-')
	}
	l = buffer.Len()
	var b []byte = make([]byte, l)
	for i := 0; i < l; i++ {
		b[i], _ = buffer.ReadByte()
	}
	left, right := 0, l-1
	for left < right {
		swapBytes(b, left, right)
		left++
		right--
	}
	return string(b)
}

func main() { //错误例子：3126.31261236
}
