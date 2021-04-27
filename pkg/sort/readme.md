## sort 包相关的核心使用方法

sort 内置的包可以很轻松的对一系列的结构进行排序，比如常见的原始数据类型的slice的排序，以及用户自定义struct 类型的slice 排序。

整体来讲，在对复杂的slice进行排序时，仅需要该slice类型实现几个方法即可： Less() int,Len(i,j int) int,Swap(i,j int) 。

同时，对于复杂的结构体来说，对某个key或者value排序，只需要适当修改一个排序函数即可。


- sort.Sort(data): 默认会根据值排序,data is a Interface(必须实现Len(),Less(i,j),Swap(i,j) 三个方法) 
- sort.Slice(x,less): 第二种根据自定义一个Less(i,j) 函数，来进行排序;当然啦x必须是一个slice，但是排序无法保证稳定性，相等的元素再次排序可能会顺序反转
- sort.SliceStable(x interface{},less func(i,j int) bool): 同上，每次排序的顺序会相对稳定
- sort.SliceIsSorted(x,less) bool: 返回一个slice是否是用less进行排序的
- sort.Strings(x []string): 对一个字符串slice 进行排序
- sort.StringsAreSorted(x) bool:
- sort.Float64s(x []float64): 直接给float64的切片进行排序
- sort.Float64sAreSorted(x) bool: 查看x是否有序
- sort.Ints(x []int): int类型排序
- sort.IntsAreSorted(x) bool: 
- sort.IsSorted(data Interface) bool: 查看data是否有序
- `sort.Reverse(data Interface) Interface`: 排序结构体进行一次反转，data必须实现了len,swap,less函数(sort内部提供了些基本变量Float64Slice,IntSlice,StringSlice)



注意: 通常情况下，我们使用sort内置的包以及自定义的结构体切片其实可以满足大部分的需求，但是可能会存在一些需要特殊的排序需求，比如需要根据整个结构体的 key 对相关业务进行排序、或者需要连续对结构体中的多个字段进行排序


### 基于排序结构的切片构造通用的排序函数(less)，用于指定多种排序函数

```

// 自定义两个新的类型
type earthMass float64
type au float64

// Planet 定义了太阳系的一些属性
type Planet struct {
	name string
  mass earthMass
	distance au
}

// 定义一个planetSorter的排序结构，可以通过封装一个排序函数,以及需要排序的结构体
// 因此排序的函数完全可以通过by函数的传递来控制
type planetSorter struct {
	planets []Planet
  by 	func(p1,p2 *Planet) bool
}

// 给planetSorter定义less,swap,len方法
func (s *planetSorter) Len() int {
	return len(s.planets)
} 
func (s *planetSorter) Swap(i,j int) {
	s.planets[i],s.planets[j] = s.planets[j],s.planets[i]
}
func (s *planetSorter) Less(i,j int) bool {
	return s.by(&s.planets[i],&s.planets[j])
}



// 抽象一个自定义的排序函数，其实就是less函数类型
type By func(p1,p2 *Planet) bool 

// 给by定义一个排序函数，默认排序参数必须是结构体的切片类型
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by: by,			// 根据初始化化时的排序函数来构造
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
	nameSort := func(p1,p2 *Planet) bool {
		return p1.name < p2.name
	}	
	massSort := func(p1,p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distanceSort := func(p1,p2 *Planet) bool {
		return p1.distance < p2.distance
	}

	// 降序排序
	decreasingDistanceSort := func(p1,p2 *Planet) bool {
		return distanceSort(p2,p2)
	}


	fmt.Println("origin metrics:",plannets)
	// 对对象排序
	By(nameSort).Sort(plannets)
	fmt.Println("nameSort:",plannets)
	By(massSort).Sort(plannets)
	fmt.Println("nameSort:",plannets)
	By(distanceSort).Sort(plannets)
	fmt.Println("nameSort:",plannets)
	By(decreasingDistanceSort).Sort(plannets)
	fmt.Println("nameSort:",plannets)

	
	
}


```


### 同时对多个key进行排序的的方法


```
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


```



