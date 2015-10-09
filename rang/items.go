package rang

import (
	"reflect"
)

func GetName() string {
	return "Hello World!"
}

type WeightedItem interface {
	GetWeight() float64
}

type Item struct {
	Weight float64
}

func NewItem(v float64) *Item {
	return &Item{v}
}

func (self Item) GetWeight() float64 {
	return self.Weight
}

func (self *Item) SetWeight(weight float64) {
	self.Weight = weight
}

func (self Item) ComparableValue() float64 {
	return self.GetWeight()
}

func (self Item) Compare(v interface{}) int16 {
	
	logic := func (other float64) int16 {
		if other == self.GetWeight() { 
			return 0 
		} else if self.GetWeight() > other { 
			return 1 
		} else { 
			return -1 
		}
	}
	
	if reflect.TypeOf(self) == reflect.TypeOf(v) {
		return logic(v.(Item).GetWeight()) 
	} else if reflect.TypeOf(&self) == reflect.TypeOf(v) {
		return logic(v.(*Item).GetWeight())		
	}
	
	return -1
}




/*func DoCompare(a, b Comparable) int32 {
	return a.Compare(b)
}*/