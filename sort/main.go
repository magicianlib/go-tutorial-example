package main

import (
	"fmt"
	"sort"
)

type People struct {
	Name   string  // 名称
	Age    uint8   // 年龄
	Height uint8   // 身高
	Weight float32 // 体重
}

type ByFunc func(p1, p2 *People) bool

func (by ByFunc) Sort(ps []People) {
	var p = &peopleSort{
		p:  ps,
		by: by,
	}

	sort.Sort(p)
}

type peopleSort struct {
	p  []People
	by ByFunc
}

func (ps *peopleSort) Len() int {
	return len(ps.p)
}

func (ps *peopleSort) Less(i, j int) bool {
	return ps.by(&ps.p[i], &ps.p[j])
}

func (ps *peopleSort) Swap(i, j int) {
	ps.p[i], ps.p[j] = ps.p[j], ps.p[i]
}

func main() {

	var peoples = []People{
		{"张三", 20, 165, 120},
		{"李四", 18, 170, 90},
		{"王二", 24, 160, 110},
	}

	ageFunc := func(p1, p2 *People) bool {
		return p1.Age < p2.Age
	}

	heightFunc := func(p1, p2 *People) bool {
		return p1.Height < p2.Height
	}

	weightFunc := func(p1, p2 *People) bool {
		return p1.Weight < p2.Weight
	}

	fmt.Println(peoples)

	ByFunc(ageFunc).Sort(peoples)
	fmt.Println("ByAge:   ", peoples)

	ByFunc(heightFunc).Sort(peoples)
	fmt.Println("ByHeight:", peoples)

	ByFunc(weightFunc).Sort(peoples)
	fmt.Println("ByWeight:", peoples)
}
