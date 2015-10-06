package main

import (
	"fmt"
	"github.com/Himself12794/learning-go/rang"
	//"math/rand"
	//"time"
)

func main() {
	testRange()
}

func testRange() {
	
	i := rang.Float64(3.14159265)
	
	a := rang.NewItem( 4 )
	b := rang.NewItem( 5 )
	//v := rang.NewItem( 70 )
	
	r := rang.RangeClosedInclusive( a, b )
	
	fmt.Println(r.Contains( i ))
	
}