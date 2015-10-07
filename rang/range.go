package rang

import (
//	"fmt"
	"strconv"
)

func DoCompare(a, b Comparable) int {
	var v int
	
	//fmt.Println("Compare value a: ", a.Value())
	//fmt.Println("Compare value b: ", b.Value())
	
	if a.Value() > b.Value() { 
		v = 1 
	} else if a.Value() == b.Value() { 
		v = 0 
	} else { 
		v = -1
	} 
	
	return v 
}

type Comparable interface {
	Value() float64
}

type Range struct {
	inclusiveA bool
	inclusiveB bool
	openA bool
	openB bool
	endA Comparable
	endB Comparable
}

	/**
	 * Return whether or not all values of range r fall within this range 
	 */
	func (self Range) ContainsRange(r Range) bool {
		return self.Contains(r.endA) && self.Contains(r.endB)
	}
	
	/**
	 * Returns whether or not the range contains the value
	 */
	func (self Range) Contains(o Comparable) bool {
		
		var a bool
		var b bool
		
		if self.openA {
			a = true
		} else if self.inclusiveA {
			a = o.Value() >= self.endA.Value()
		} else {
			a = o.Value() > self.endA.Value()
		}
		
		if self.openB {
			b = true
		} else if self.inclusiveB {
			b = o.Value() <= self.endB.Value()
		} else {
			b = o.Value() < self.endB.Value()
		}
		
		return a && b
	}
	
	func (self Range) String() string {
		return "End A: " + 
			strconv.FormatFloat(self.endA.Value(), 'f', -1, 64) + 
			", End B: " + 
			strconv.FormatFloat(self.endB.Value(), 'f', -1, 64)
	}

/**
 * Range of the form a < n < b 
 */
func RangeClosedExclusive(a, b Comparable) Range {
	return Range{inclusiveA: false, inclusiveB: false, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a <= n <= b 
 */
func RangeClosedInclusive(a, b Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a <= n < b 
 */
func RangeClosedInclusiveLeft(a, b Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: false, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a < n <= b 
 */
func RangeClosedInclusiveRight(a, b Comparable) Range {
	return Range{inclusiveA: false, inclusiveB: true, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range containing all elements
 */
func RangeAll() Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: true, openB: true, endA: nil, endB: nil}
}

/**
 * Range of the form a <= n 
 */
func RangeOpenLeftInclusive(a Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: true, endA: a, endB: nil}
}

/**
 * Range of the form a < n
 */
func RangeOpenLeftExclusive(a Comparable) Range {
	return Range{inclusiveA: false, inclusiveB: true, openA: false, openB: true, endA: a, endB: nil}
}

/**
 * Range of the form n <= b 
 */
func RangeOpenRightInclusive(a Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: true, openB: false, endA: nil, endB: a}
}

/**
 * Range of the form n < b
 */
func RangeOpenRightExclusive(a Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: false, openA: true, openB: false, endA: nil, endB: a}
}

/**
 * Range Containing a single value
 */
func RangeSingleton(a Comparable) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: false, endA: a, endB: a}
}

