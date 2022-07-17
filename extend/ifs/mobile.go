package main

import "fmt"

type Brand interface {
	GetName() string
}

func PrintBrandName(b Brand) {
	fmt.Println(b.GetName())
}

type Huawei struct {
	Name string
}

func (h Huawei) GetName() string {
	return h.Name
}

type iPhone struct {
	Name string
}

func (i iPhone) GetName() string {
	return i.Name
}

type iPhoneSE struct {
	Name string
	iPhone
}

// func (i iPhoneSE) GetName() string {
// 	return i.Name
// }

func main() {
	var h = Huawei{Name: "Huawei"}
	PrintBrandName(h)

	var i = iPhone{Name: "iPhone"}
	PrintBrandName(i)

	var se = iPhoneSE{
		Name:   "iPhoneSE",
		iPhone: iPhone{Name: "iPhone"},
	}
	PrintBrandName(se)
}
