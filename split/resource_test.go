package split

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	// 硬性要求：必须大于guarantee 小于 capacity
	// 软性要求：必须尽可能贴和weight 和 request
	// 如果request = 0, 则weight 视为 0，表示不参与分配

	try_num := 1000
	queue_num := 4

	total := &Resource{
		MilliCPU: 100,
		Memory:   100,
		//ScalarResources: map[string]float64{
		//	"gpu": 8,
		//},
	}
	weights := []int{1, 2, 3, 4} // weight * total 后的值
	guaranteeCpu := []float64{20, 0, 10, 0}
	guaranteeMem := []float64{20, 0, 10, 0}
	guarantees := Resources(guaranteeCpu, guaranteeMem)
	capacityCpu := []float64{30, 60, 70, 10}
	capacityMem := []float64{30, 60, 70, 10}
	capacities := Resources(capacityCpu, capacityMem)
	requestCpu := []float64{30, 60, 0, 10}
	requestMem := []float64{30, 60, 0, 10}
	requests := Resources(requestCpu, requestMem)

	Cal(total, guarantees, capacities, requests, weights, queue_num, try_num)

	//[cpu 27.20, memory 23.40 cpu 43.50, memory 35.40 cpu 12.90, memory 19.40 cpu 16.40, memory 21.80]
}

func Resources(cpus, mems []float64) []*Resource {
	resources := make([]*Resource, 0)
	for i := 0; i < len(cpus); i++ {
		resource := &Resource{
			MilliCPU: cpus[i],
			Memory:   mems[i],
		}
		resources = append(resources, resource)
	}
	return resources
}

func TestRandomFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomFloat64(100, 1))
	}
}

func TestRandomNumFloat64(t *testing.T) {
	data := RandomNumFloat64(100, 4, 1)
	fmt.Println(data)
}
