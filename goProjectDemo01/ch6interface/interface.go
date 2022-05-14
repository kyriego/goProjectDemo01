package main

import (
	"fmt"
	"io"
)

//string,bool,int,uint,float32,float64
func foo(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "null"
	case string:
		return x
	case bool:
		if x {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case int:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case float32:
		return fmt.Sprintf("%f", x)
	case float64:
		return fmt.Sprintf("%f", x)
	default:
		panic(fmt.Sprintf("unexpected type!"))
	}
}

func foo1(w io.Writer, str string) {
	type stringWriter interface {
		WriteString(str string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		sw.WriteString(str)
	} else {
		w.Write([]byte(str))
	}
}

//string,bool,int,uint,float32,float64
func foo2(x interface{}) string {
	if s, ok := x.(string); ok {
		return s
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		} else {
			return "FALSE"
		}
	} else if i, ok := x.(int); ok {
		return fmt.Sprintf("%d", i)
	} else if u, ok := x.(uint); ok {
		return fmt.Sprintf("%d", u)
	} else if f, ok := x.(float32); ok {
		return fmt.Sprintf("%f", f)
	} else if f2, ok := x.(float64); ok {
		return fmt.Sprintf("%f", f2)
	} else {
		panic(fmt.Sprintf("unexpected type!"))
	}
}

func main() {
	fmt.Printf("foo2(false): %v\n", foo2(false))
}
