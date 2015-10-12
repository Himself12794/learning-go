package main

import (
	"fmt"
	//"sort"
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
	
	keys := make([]rang.Comparable, len(valMap))
	
	i := 0
	for k, v := range valMap {
		keys[i] = k
		fmt.Println("Value:", k.GetWeight(), "was chosen:", v, "times")
	} 
	
	rang.InsertionSort(keys)
	
	for _, v := range keys {
		fmt.Println(v)
	}
}