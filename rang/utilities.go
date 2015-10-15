package rang

import (
	"math/rand"
	"time"
)


var randGenerator *rand.Rand = rand.New( rand.NewSource( time.Now().Unix() ))

func SelectRandomWeightedItem(items []WeightedItem) WeightedItem {
	ranges := make(map[Range]WeightedItem)
	
	total := 0.0
	for i := 0; i < len(items); i++{
		curr := items[i]
		
		if curr.GetWeight() > 0 {
			
			tmpWeight := total + curr.GetWeight()
			tmp := RangeClosedInclusiveLeft( Float64(total), Float64(tmpWeight) )
			total = tmpWeight
			ranges[tmp] = curr

		}
	}
	
	var choice float64
	
	randValue := randGenerator.Float64()
	choice = randValue * total
	
	for k, v := range ranges {
		
		if k.Contains(Float64(choice)){
			return v
		}
		
	}
	
	return nil
}

func InsertionSort(items []Comparable) {
	
	 for i := 1; i < len(items) - 1; i++ {
    	x := items[i]
    	var j int
    	for j = i - 1; j >= 0 && x.Compare(items[j]) <= -1; j--  {
       		items[j + 1] = items[j]
    	}
    	items[j + 1] = x
	 }
	
}
