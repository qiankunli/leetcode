package split

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func Cal(total *Resource, guarantees, capacities, requests []*Resource, weights []int, queue_num, try_num int) {

	/**
	1. 计算weight_resources
	2. 先确保 各个queue 得到guarantee，剩下的remaining 用于在各个queue 二次分配
	3. 将 remaining 随机分为 queue_num 份分给每个queue，并保证 每个queue的resource 在 [guarantee,capacity] 范围内
		1. 将remaining 随机分成 queue_num 份 分给每个queue
		2. 每个queue 此时的资源 = 随机从remaining 分得的一份 split + guarantee，split + guarantee可能超过queue的capacity
		3. 如果超过，则将多出来的部分 extra = split + guarantee - capacity，匀给下一个queue，知道最后一个queue
	4. 计算本次随机分配的cost，cost 由 costw 和 costr 组成
		1. costw 表示 本次分配与 weight_resources 的距离，数学上为 两个张量的距离，有多种计算策略，示例中使用欧拉距离
		2. costr 表示 本次分配与 request 的距离
		3. cost = ww * costw + wr * costr， 通过ww 和 wr的值，可以控制更希望 贴合weight_resources 还是request
	5. 重复第3,4 步 try_num次，记录最小成本时的分配方案，即为最终分配方案。try_num 次数越多，效果越好。
	*/
	// 根据weight 比例 计算每个queue 对应的 weight_resources
	weight_sum := sum(weights)
	weight_resources := make([]*Resource, 0)
	for _, w := range weights {
		weight_resources = append(weight_resources, total.Clone().Multi(float64(w)/float64(weight_sum)))
	}

	total_guarantees := EmptyResource()
	for queue := 0; queue < queue_num; queue++ {
		total_guarantees.Add(guarantees[queue])
	}
	// 计算 remaining
	remaining := total.Clone().sub(total_guarantees)
	min := math.MaxFloat64
	var d []*Resource
	// 找到 try_num 次随机分配方案中 cost 最小的方案
	for i := 0; i < try_num; i++ {
		// 一次随机分配
		deserves := try(remaining.Clone(), guarantees, capacities, queue_num)
		// 计算此次随机分配的cost
		costw := cost(deserves, weight_resources)
		costr := cost(deserves, requests)
		cost := costw + 2*costr
		if min > cost {
			min = cost
			d = deserves
		}
	}
	fmt.Println(d)

}

// 将 remaining 随机分为 queue_num 份分给每个queue，并保证 每个queue的resource 在 [guarantee,capacity] 范围内
func try(remaining *Resource, guarantees, capacities []*Resource, queue_num int) []*Resource {
	deserves := make([]*Resource, 0)
	splits := split(remaining, queue_num)
	deserve := EmptyResource()
	extra := EmptyResource()
	for queue := 0; queue < queue_num; queue++ {
		// 把guarantee 先减了，且保证sum=total
		deserve, extra = mix(extra, splits[queue], guarantees[queue], capacities[queue])
		deserves = append(deserves, deserve)
	}
	return deserves
}

// 将超过 capacity 的资源匀出来，留着给下一个queue用
func mix(extra, split, guarantee, capacity *Resource) (*Resource, *Resource) {
	deserve := extra.Add(split).Add(guarantee)
	if deserve.LessEqual(capacity, Infinity) {
		return deserve, EmptyResource()
	}
	d := deserve.MinDimensionResource(capacity, Infinity)
	return d, deserve.Sub(d)
}

// 将remaining 随机分成 queue_num 份儿
func split(remaining *Resource, queue_num int) []*Resource {
	result := make([]*Resource, 0)
	cpuSplit := RandomNumFloat64(remaining.MilliCPU, queue_num, 1)
	memSplit := RandomNumFloat64(remaining.Memory, queue_num, 1)
	scalarSplit := make(map[string][]float64)
	if len(remaining.ScalarResources) > 0 {
		for r, v := range remaining.ScalarResources {
			rSplit := RandomNumFloat64(v, queue_num, 1)
			scalarSplit[r] = rSplit
		}
	}

	for i := 0; i < queue_num; i++ {
		q := &Resource{
			MilliCPU:        cpuSplit[i],
			Memory:          memSplit[i],
			ScalarResources: make(map[string]float64),
		}
		result = append(result, q)
		if len(remaining.ScalarResources) == 0 {
			continue
		}
		for r, v := range scalarSplit {
			q.ScalarResources[r] = v[i]
		}
	}
	return result
}

func cost(l, r []*Resource) float64 {
	cost := 0.0
	for i := 0; i < len(l); i++ {
		cost += l[i].Cost(r[i])
	}
	return cost
}

func (r *Resource) Cost(rr *Resource) float64 {
	costCpu := math.Pow(r.MilliCPU-rr.MilliCPU, 2)
	costMem := math.Pow(r.Memory-rr.Memory, 2)
	cost := costCpu + costMem
	for rName, rQuant := range rr.ScalarResources {
		if r.ScalarResources == nil {
			r.ScalarResources = map[string]float64{}
		}
		cost += math.Pow(r.ScalarResources[rName]-rQuant, 2)
	}
	return cost
}

func min(l, r float64) float64 {
	if l > r {
		return r
	}
	return l
}
func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func RandomFloat64(max float64, decimal int) float64 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()
	random = random * max
	pow := math.Pow10(decimal)
	randomInt64 := int64(random * pow)
	randomFloat64 := float64(randomInt64) / pow
	return randomFloat64
}

// 截绳法
func RandomNumFloat64(sum float64, num int, decimal int) []float64 {
	if num <= 0 {
		return nil
	}
	if num == 1 {
		return []float64{RandomFloat64(sum, decimal)}
	}
	nodes := make([]float64, 0)
	nodes = append(nodes, 0.0)
	for i := 0; i < num-1; i++ {
		node := RandomFloat64(sum, decimal)
		nodes = append(nodes, node)
	}
	nodes = append(nodes, sum)
	sort.Float64s(nodes)
	result := make([]float64, 0)
	for i := 1; i < len(nodes); i++ {
		result = append(result, sub(nodes[i], nodes[i-1], decimal))
	}
	return result
}

func sub(l, r float64, decimal int) float64 {
	pow := math.Pow10(decimal)
	subInt64 := int64(l*pow) - int64(r*pow)
	subFloat64 := float64(subInt64) / pow
	return subFloat64
}
