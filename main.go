package main

import (
	"fmt"
	//"sort"
	"github.com/Himself12794/learning-go/rang"
	loc "github.com/Himself12794/learning-go/location"
	oauth "github.com/go-oauth2/oauth2"
)

func main() {
	//testRange()
	testLocation()
	testFunc()
	fmt.PrintF(oauth.Code)
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
	
	fmt.Println( a.Compare( b ) )
}

func testFunc() {
	f := func (t rang.Float64) {
		fmt.Println(t.ComparableValue())
	}
	
	f(6.5)
}

func testLocation() {
	v1 := loc.NewVec3(3, 4, 0)
	v2 := loc.NewVec3(1, 1, 1)
	v3 := loc.NewVec3(0.54, 10.44488522, 0.128872132)
	
	fmt.Println(v1.DotProduct(v2))
	fmt.Println(v1.Normalize())
	fmt.Println(v3.CrossProduct(v2))
	

}
