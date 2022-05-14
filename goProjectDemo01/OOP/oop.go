package main

import (
	"errors"
	"fmt"
)

//tips: 思考一下map[string][]string的插入和查找逻辑

//OOP一.对象的创建（单个：基本数据类型，string，数组，slice，map。 组合：结构体）
//1.决定方法接收者是一个普通变量还是指针时，先考虑对象的数据类型进行参数传递的时候是引用传递还是值传递，再决定是否用指针
//换句话说，如果本来就是指针，则不用考虑（slice的话，因为更新的时候需要对slice引用更新例如slice = append(slice, num)，虽然传递的是引用，但会导致外面的内部的两个slice不一致，所以需要修改的话还是要用指针）
//，如果不是指针，那么就可以根据是否需要修改来决定是否用指针作为接收者
//2.对象的创建： ①结构体变量， ②单个变量的命名类型就用原来对应类型的字面量，只是换了一个名字  [例子：封装一个数组对象，map[string][]string的封装]
type MyMap map[string][]string

func (m MyMap) Add(key, val string) error {
	if m == nil {
		return errors.New("can not add entry to a nil map!")
	}
	if _, ok := m[key]; !ok {
		m[key] = make([]string, 0)
	}
	m[key] = append(m[key], val)
	return nil
}

func (m MyMap) Add1(key, val string) error {
	if m == nil {
		return errors.New("can not add entry to a nil map!")
	}
	if _, ok := m[key]; !ok {
		m[key] = make([]string, 0)
	}
	m[key] = append(m[key], val)
	return nil
}

func (m MyMap) Get(key string) string {
	if m == nil {
		return ""
	}
	strs := m[key]
	if len(strs) == 0 {
		return ""
	}
	return strs[0]
}

//数据结构  map[string][]string
func (m MyMap) Gett(key string) string {
	if m == nil {
		return ""
	}
	if strs := m[key]; len(strs) > 0 {
		return strs[0]
	}
	return ""
}

type MySlice []int

func NewMySlice(args ...int) MySlice {
	ms := MySlice(make([]int, 0))
	for _, n := range args {
		ms = append(ms, n)
	}
	return ms
}

func (m *MySlice) Add(val int) {
	*m = append(*m, val)
}

func (m MySlice) Remove() {
	m = m[:len(m)-1]
}

func (m MySlice) Print() {
	if len(m) < 0 {
		return
	}
	for i, v := range m {
		fmt.Print(v)
		if i != len(m)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}

type MyMap1 map[string]string

func (m MyMap1) Add(key, val string) error {
	if m == nil {
		return errors.New("blablabla")
	}
	m[key] = val
	return nil
}

func (m MyMap1) Get(key string) string {
	if m == nil {
		return ""
	}
	return m[key]
}

func (m MyMap1) Print() {
	for k, v := range m {
		fmt.Printf("%s:%s\n", k, v)
	}
}

func main() {
	/* 	m := map[string]string{
		"kyrieGo": "nets",
		"james":   "laker",
		"curry":   "warrior",
	} */
	var nums []int
	nums = append(nums, 1)
	fmt.Print(nums)
}
