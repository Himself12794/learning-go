package rang

import (
)

func GetName() string {
	return "Hello World!"
}

type WeightedItem interface {
	SetWeight(w float64)
	GetWeight() float64
}

type Item struct {
	weight float64
}

func NewItem(v float64) *Item {
	return &Item{v}
}

func (self Item) GetWeight() float64 {
	return self.weight
}

func (self *Item) SetWeight(weight float64) {
	self.weight = weight
}

func (self Item) Value() float64 {
	return self.GetWeight()
}


/*func DoCompare(a, b Comparable) int32 {
	return a.Compare(b)
}*/