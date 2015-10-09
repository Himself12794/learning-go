package main

import (
	"fmt"
	"github.com/Himself12794/learning-go/rang"
)

func main() {
	testRange()
}

func testRange() {
	
	a := rang.NewItem( 10 )
	b := rang.NewItem( 1 )
	ar := []rang.WeightedItem{ a, b }
	valMap := make(map[*rang.Item]int)
	
	var c *rang.Item
	for i := 0; i < 500; i++ {
		c = rang.SelectRandomWeightedItem( ar ).(*rang.Item)
		valMap[c] += 1
	}
	
	for k, v := range valMap {
		fmt.Println("Value:", k.GetWeight(), "was chosen:", v, "times")
	} 
	
	fmt.Println( a.Compare( b ) )
}