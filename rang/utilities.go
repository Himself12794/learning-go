package rang

import (
	"math/rand"
	"time"
)


var randGenerator *rand.Rand = rand.New( rand.NewSource( time.Now().Unix() ))

func SelectRandomWeightedItem(items []WeightedItem) WeightedItem {
	ranges := make(map[Range]WeightedItem)
	ranges[RangeSingleton( Float64(50) )] = items[0]
	
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
	
	randValue := randGenerator.Float64()
	choice := randValue * total
	
	for k, v := range ranges {
		
		if k.Contains(Float64(choice)){
			return v
		}
		
	}
	
	return nil
}