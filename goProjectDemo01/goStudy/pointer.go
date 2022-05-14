package main

import (
	"flag"
	"fmt"
	"strings"
)

//熟悉一下指针的使用：
//案例1：创建两个int变量，通过指针来交换它们之间的值
//案例2：定义Person结构体类型，创建一个结构体变量，编写一个函数，函数接收一个结构体变量指针来进行修改
//(基本的bean类的编写：结构体可导出，成员变量私有，所有的构造，getter/setter方法可导出)
//案例3：创建一个int变量，编写一个函数来接收该int变量地址，然后在函数中修改这个变量的值
//案例4：对echo进行升级，定义两个终端参数，n用来控制输出结果是否换行，seq用来定义参数们之间的分割符号
//案例5：new关键字的使用： new(type)：创建一个type类型变量并且返回其地址/说白了就是一个指针——>延伸至对象的构造函数
//积累：运行go程序制定参数时要先输入“命名参数”（记得加逗号）然后就是“无命名参数”   读取的时候先flag.Parse，然后flag.Args来读取“无命名参数”,下标从0开始
var n = flag.Bool("n", false, "decide if Line feed")
var seq = flag.String("seq", "", "the seq of args")

func echo1() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		str := strings.Join(args, *seq)
		fmt.Print(str)
		if *n {
			fmt.Println()
		}
	} else {
		return
	}
}

func add(num *int) {
	(*num)++
}

func main() {
	var a int = 100
	add(&a)
	fmt.Print(a)
}
