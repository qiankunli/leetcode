/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package split

import (
	"fmt"
	"math"
)

const (
	// GPUResourceName need to follow https://github.com/NVIDIA/k8s-device-plugin/blob/66a35b71ac4b5cbfb04714678b548bd77e5ba719/server.go#L20
	GPUResourceName = "nvidia.com/gpu"
)

const (
	minResource float64 = 0.1
)

// DimensionDefaultValue means default value for black resource dimension
type DimensionDefaultValue int

const (
	// Zero means resource dimension not defined will be treated as zero
	Zero DimensionDefaultValue = 0
	// Infinity means resource dimension not defined will be treated as infinity
	Infinity DimensionDefaultValue = -1
)

// Resource struct defines all the resource type
type Resource struct {
	MilliCPU float64
	Memory   float64

	// ScalarResources
	ScalarResources map[string]float64

	// MaxTaskNum is only used by predicates; it should NOT
	// be accounted in other operators, e.g. Add.
	MaxTaskNum int
}

// EmptyResource creates a empty resource object and returns
func EmptyResource() *Resource {
	return &Resource{}
}

// Clone is used to clone a resource type, which is a deep copy function.
func (r *Resource) Clone() *Resource {
	clone := &Resource{
		MilliCPU:   r.MilliCPU,
		Memory:     r.Memory,
		MaxTaskNum: r.MaxTaskNum,
	}

	if r.ScalarResources != nil {
		clone.ScalarResources = make(map[string]float64)
		for k, v := range r.ScalarResources {
			clone.ScalarResources[k] = v
		}
	}

	return clone
}

// String returns resource details in string format
func (r *Resource) String() string {
	str := fmt.Sprintf("cpu %0.2f, memory %0.2f", r.MilliCPU, r.Memory)
	for rName, rQuant := range r.ScalarResources {
		str = fmt.Sprintf("%s, %s %0.2f", str, rName, rQuant)
	}
	return str
}

// Get returns the resource value for that particular resource type
func (r *Resource) Get(rn string) float64 {
	switch rn {
	case "cpu":
		return r.MilliCPU
	case "memory":
		return r.Memory
	default:
		if r.ScalarResources == nil {
			return 0
		}
		return r.ScalarResources[rn]
	}
}

// IsEmpty returns false if any kind of resource is not less than min value, otherwise returns true
func (r *Resource) IsEmpty() bool {
	if !(r.MilliCPU < minResource && r.Memory < minResource) {
		return false
	}

	for _, rQuant := range r.ScalarResources {
		if rQuant >= minResource {
			return false
		}
	}

	return true
}

// Add is used to add two given resources
func (r *Resource) Add(rr *Resource) *Resource {
	r.MilliCPU += rr.MilliCPU
	r.Memory += rr.Memory

	for rName, rQuant := range rr.ScalarResources {
		if r.ScalarResources == nil {
			r.ScalarResources = map[string]float64{}
		}
		r.ScalarResources[rName] += rQuant
	}

	return r
}

// Add is used to add two given resources
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

// Sub subtracts two Resource objects with assertion.
func (r *Resource) Sub(rr *Resource) *Resource {
	return r.sub(rr)
}

// sub subtracts two Resource objects.
func (r *Resource) sub(rr *Resource) *Resource {
	r.MilliCPU -= rr.MilliCPU
	r.Memory -= rr.Memory

	if r.ScalarResources == nil {
		return r
	}
	for rrName, rrQuant := range rr.ScalarResources {
		r.ScalarResources[rrName] -= rrQuant
	}

	return r
}

// Multi multiples the resource with ratio provided
func (r *Resource) Multi(ratio float64) *Resource {
	r.MilliCPU *= ratio
	r.Memory *= ratio
	for rName, rQuant := range r.ScalarResources {
		r.ScalarResources[rName] = rQuant * ratio
	}
	return r
}

// SetMaxResource compares with ResourceList and takes max value for each Resource.
func (r *Resource) SetMaxResource(rr *Resource) {
	if r == nil || rr == nil {
		return
	}

	if rr.MilliCPU > r.MilliCPU {
		r.MilliCPU = rr.MilliCPU
	}
	if rr.Memory > r.Memory {
		r.Memory = rr.Memory
	}

	for rrName, rrQuant := range rr.ScalarResources {
		if r.ScalarResources == nil {
			r.ScalarResources = make(map[string]float64)
			for k, v := range rr.ScalarResources {
				r.ScalarResources[k] = v
			}
			return
		}
		_, ok := r.ScalarResources[rrName]
		if !ok || rrQuant > r.ScalarResources[rrName] {
			r.ScalarResources[rrName] = rrQuant
		}
	}
}
func (r *Resource) MinDimensionResource(rr *Resource, defaultValue DimensionDefaultValue) *Resource {
	if rr.MilliCPU < r.MilliCPU {
		r.MilliCPU = rr.MilliCPU
	}
	if rr.Memory < r.Memory {
		r.Memory = rr.Memory
	}

	if r.ScalarResources == nil {
		return r
	}

	if rr.ScalarResources == nil {
		if defaultValue == Infinity {
			return r
		}

		for name := range r.ScalarResources {
			r.ScalarResources[name] = 0
		}
		return r
	}

	for name, quant := range r.ScalarResources {
		rQuant, ok := rr.ScalarResources[name]
		if ok {
			r.ScalarResources[name] = math.Min(quant, rQuant)
		} else {
			if defaultValue == Infinity {
				continue
			}

			r.ScalarResources[name] = 0
		}
	}
	return r
}
func (r *Resource) LessEqual(rr *Resource, defaultValue DimensionDefaultValue) bool {
	lessEqualFunc := func(l, r, diff float64) bool {
		if l < r || math.Abs(l-r) < diff {
			return true
		}
		return false
	}

	if !lessEqualFunc(r.MilliCPU, rr.MilliCPU, minResource) {
		return false
	}
	if !lessEqualFunc(r.Memory, rr.Memory, minResource) {
		return false
	}

	for resourceName, leftValue := range r.ScalarResources {
		rightValue, ok := rr.ScalarResources[resourceName]
		if !ok && defaultValue == Infinity {
			continue
		}

		if !lessEqualFunc(leftValue, rightValue, minResource) {
			return false
		}
	}
	return true
}
