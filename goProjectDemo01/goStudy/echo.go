package main //打印所有的命令行参数
import (
	"fmt"
	"os"
	"strings"
)

//①使用索引Index的方式i来进行遍历
func argsToString1() string {
	var res string
	args := os.Args
	n := len(args)
	for i := 1; i < len(args); i++ {
		res += os.Args[i]
		if i != n-1 {
			res += " "
		}
	}
	return res
}

//②使用range的方式来遍历Args

func argsToString2() string {
	var res string
	args := os.Args[1:]
	for _, arg := range args {
		res += arg + " "
	}
	return res
}

func argsToString3() string {
	args := os.Args[1:]
	res := strings.Join(args, " ")
	return res
}

func main() {
	fmt.Println(argsToString3())
}
