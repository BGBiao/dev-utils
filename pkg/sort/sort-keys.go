package main

import (
	"fmt"
	"sort"
)

// 自定义两个新的类型
type earthMass float64
type au float64

// Planet 定义了太阳系的一些属性
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

// 定义一个planetSorter的排序结构，可以通过封装一个排序函数,以及需要排序的结构体
// 因此排序的函数完全可以通过by函数的传递来控制
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool
}

// 给planetSorter定义less,swap,len方法
func (s *planetSorter) Len() int {
	return len(s.planets)
}
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

// 抽象一个自定义的排序函数，其实就是less函数类型
type By func(p1, p2 *Planet) bool

// 给by定义一个排序函数，默认排序参数必须是结构体的切片类型
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // 根据初始化化时的排序函数来构造
	}

	sort.Sort(ps)
}

func main() {
	var planets = []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}

	// 定义一个By 的类型，实际是Less类函数
	// 默认都是使用< 也就是升序排序
	nameSort := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	massSort := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distanceSort := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}

	// 降序排序
	decreasingDistanceSort := func(p1, p2 *Planet) bool {
		return distanceSort(p2, p1)
	}

	fmt.Println("origin metrics:", planets)
	// 对对象排序
	By(nameSort).Sort(planets)
	fmt.Println("nameSort:", planets)
	By(massSort).Sort(planets)
	fmt.Println("massSort:", planets)
	By(distanceSort).Sort(planets)
	fmt.Println("distanceSort:", planets)
	By(decreasingDistanceSort).Sort(planets)
	fmt.Println("decreasingDistanceSort:", planets)

}
