package main

import (
	"fmt"
	"sort"
)

// 原始待排序的结构体
type ITer struct {
	name     string
	position string
	language string
	age      int
}

// 定义通用的排序函数lessFunc
type lessFunc func(p1, p2 *ITer) bool

// 封装一个排序类型
// 该结构体其实默认就实现了Sort 的接口了(less,swap,len)
type multiSorter struct {
	iters []ITer     // 排序切片
	less  []lessFunc // 排序函数切片
}

func (ms *multiSorter) Len() int {
	return len(ms.iters)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.iters[i], ms.iters[j] = ms.iters[j], ms.iters[i]
}

// 沿着less函数来挨个遍历，直到找到两者中间最小的
// 其实没太看懂
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.iters[i], &ms.iters[j]
	// 开始遍历全部的排序函数
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}

	return ms.less[k](p, q)
}

// 封装一个排序方法
func (ms *multiSorter) Sort(iters []ITer) {
	ms.iters = iters
	sort.Sort(ms)
}

// 构造多排序函数
func OrderBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}

}

func main() {
	name := func(it1, it2 *ITer) bool {
		return it1.name > it2.name
	}

	position := func(it1, it2 *ITer) bool {
		return it1.position > it2.position
	}

	language := func(it1, it2 *ITer) bool {
		return it1.language > it2.language
	}

	age := func(it1, it2 *ITer) bool {
		return it1.age > it2.age
	}

	var iters = []ITer{
		{"bgbiao", "sport", "golang", 21},
		{"xxb", "ps4", "python", 28},
		{"biao", "switch", "java", 26},
		{"weichuangxxb", "it", "vue", 31},
		{"yu", "point", "html", 29},
		{"maomao", "eat", "golang+java+python", 35},
	}

	fmt.Printf("origin data:%v\n", iters)

	OrderBy(name, age).Sort(iters)
	fmt.Printf("by name and age:%v\n", iters)

	OrderBy(name, language).Sort(iters)
	fmt.Printf("by name and language:%v\n", iters)

	OrderBy(name, position, language, age).Sort(iters)
	fmt.Printf("by name,position.language,age:%v\n", iters)

	OrderBy(age).Sort(iters)
	fmt.Printf("by age:%v\n", iters)

	OrderBy(name).Sort(iters)
	fmt.Printf("by name:%v\n", iters)
}
