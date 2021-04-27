package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Metric struct {
	Id      string
	Ip      string
	Version string
	Max     float64
	Min     float64
	Avg     float64
	Current float64
}

// 定义结构体的slice 后，使用Less 对其Current指标进行降序排序
type Metrics []Metric

func (ms Metrics) Len() int { return len(ms) }
func (ms Metrics) Less(i, j int) bool {
	return ms[i].Current > ms[j].Current
}
func (ms Metrics) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	var metrics Metrics

	for i := 0; i < 10; i++ {
		tmpMetric := Metric{
			Id:      "11111",
			Ip:      "1.1.1.1",
			Max:     1.00 + rand.Float64(),
			Min:     2.00 + rand.Float64(),
			Avg:     3.00 + rand.Float64(),
			Current: 4.00 + rand.Float64(),
		}

		metrics = append(metrics, tmpMetric)
	}

	fmt.Printf("原始metrics:\n%v\n", metrics)
	// 根据自定义的less进行顺序排序
	sort.Sort(metrics)
	fmt.Printf("默认Current指标降序排序metrics:\n%v\n", metrics)
	// 自定义排序字段 使用Max指标进行降序排序
	sort.SliceStable(metrics, func(i, j int) bool {
		return metrics[i].Max > metrics[j].Max
	})
	fmt.Printf("自定义Max指标降序排序metrics:\n%v\n", metrics)

}
