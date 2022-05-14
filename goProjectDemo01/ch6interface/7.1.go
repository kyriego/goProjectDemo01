package main

import (
	"bufio"
	"flag"
	"io"
)

//7.1
//记录1: 面向对象设计： 给类型绑定方法的时候，须制定绑定到T还是绑定到*T,然后在调用的时候
//来决定是通过变量还是通过变量指针来进行调用
type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	res := 0
	if len(p) == 0 {
		return 0, nil
	}
	for i := 0; i < len(p); {
		advance, _, _ := bufio.ScanWords(p[i:], true)
		*w++
		res++
		i += advance
	}
	return res, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	res := 0
	for i := 0; i < len(p); {
		advance, _, _ := bufio.ScanLines(p[i:], true)
		*l++
		res++
		i += advance
	}
	return res, nil
}

//7.4
type MyStringReader struct {
	val   string
	index int
}

func NewMyStringReader(val string) *MyStringReader {
	var myStringReader MyStringReader
	myStringReader.val = val
	myStringReader.index = 0
	return &myStringReader
}

func (m *MyStringReader) Read(p []byte) (int, error) {
	if m.index >= len(m.val) {
		return 0, nil
	}
	n := copy(p, []byte(m.val)[m.index:])
	m.index += n
	if n < len(p) {
		return n, io.EOF
	} else {
		return n, nil
	}
}

func NewReader(str string) io.Reader {
	return NewMyStringReader(str)
}

type MyLimitReader struct {
	R io.Reader
	N int
}

func NewMyLimitReader(r io.Reader, n int) *MyLimitReader {
	myLimitReader := new(MyLimitReader)
	myLimitReader.R = r
	myLimitReader.N = n
	return myLimitReader
}

func (m *MyLimitReader) Read(p []byte) (int, error) {
	if m.N <= 0 {
		return 0, io.EOF
	}
	if len(p) > m.N {
		p = p[0:m.N]
	}
	n, err := m.R.Read(p)
	m.N -= n
	return n, err
}

func main() {
	//7.1Demo
	/* 	f, _ := os.Open("protoBuf使用记录.txt")
	   	bytes, _ := ioutil.ReadAll(f)
	   	var wc WordCounter
	   	var lc LineCounter
	   	n, _ := wc.Write(bytes)
	   	l, _ := lc.Write(bytes)
	   	lc.Write(bytes)
	   	fmt.Fprint(&lc, string(bytes))
	   	fmt.Printf("write:%d\tcounter:%d\n", l, lc) */
	//7.2Demo

	//7.4Demo
	/* 	r := MyStringsNewReader("this is the message from client")
	   	var bytes []byte = make([]byte, len(*r))
	   	r.Read(bytes)
	   	fmt.Printf("%s\n", bytes) */
	/* 	f, _ := os.Open("protoBuf使用记录.txt")
	   	var bytes []byte = make([]byte, 1024)
	   	f.Read(bytes)
	   	fmt.Printf("%s\n", bytes)
	   } */
	/* 	var bytes []byte = make([]byte, 2)
	   	reader := NewReader("this is the message from client")
	   	for {
	   		n, err := reader.Read(bytes)
	   		if err != nil {
	   			if err == io.EOF {
	   				fmt.Printf("数据已读取完毕!")
	   				break
	   			} else {
	   				fmt.Printf("读取数据出错!")
	   				break
	   			}
	   		}
	   		fmt.Printf("%s\n", bytes[:n])
	   	} */
	/* 	mlr := NewMyLimitReader(os.Stdin, 10)
	   	var bytes []byte = make([]byte, 20)
	   	mlr.Read(bytes)
	   	fmt.Printf("%s\n", bytes) */
	flag.CommandLine.IntVar()
}
