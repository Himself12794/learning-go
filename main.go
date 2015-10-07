package main

import (
	"fmt"
	"github.com/Himself12794/learning-go/rang"
)

func main() {
	testRange()
}

func testRange() {
	
	a := rang.NewItem( 5 )
	b := rang.NewItem( 0.5 )
	ar := []rang.WeightedItem{ a, b }
	
	am := 1
	var c rang.WeightedItem
	for true {
		c = rang.SelectRandomWeightedItem( ar )
		if c.GetWeight() == b.GetWeight() {
			break
		}
		am++
	}
	fmt.Println("Choice:", c.GetWeight(), "- Iterations to be chosen:", am)
}