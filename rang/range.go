package rang

import (
	"strconv"
)

func DoCompare(a, b ComparableR) int {
	var v int
	
	//fmt.Println("Compare value a: ", a.Value())
	//fmt.Println("Compare value b: ", b.Value())
	
	if a.ComparableValue() > b.ComparableValue() { 
		v = 1 
	} else if a.ComparableValue() == b.ComparableValue() { 
		v = 0 
	} else { 
		v = -1
	} 
	
	return v 
}

type Comparable interface {
	Compare(v interface{}) int16
}

// Due to no parameterized types, we must require the range ends to
// provide a value for comparison
type ComparableR interface {
	ComparableValue() float64
}

type Range struct {
	inclusiveA bool
	inclusiveB bool
	openA bool
	openB bool
	endA ComparableR
	endB ComparableR
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
	func (self Range) Contains(o ComparableR) bool {
		
		var a bool
		var b bool
		
		if self.openA {
			a = true
		} else if self.inclusiveA {
			a = o.ComparableValue() >= self.endA.ComparableValue()
		} else {
			a = o.ComparableValue() > self.endA.ComparableValue()
		}
		
		if self.openB {
			b = true
		} else if self.inclusiveB {
			b = o.ComparableValue() <= self.endB.ComparableValue()
		} else {
			b = o.ComparableValue() < self.endB.ComparableValue()
		}
		
		return a && b
	}
	
	func (self Range) String() string {
		return "End A: " + 
			strconv.FormatFloat(self.endA.ComparableValue(), 'f', -1, 64) + 
			", End B: " + 
			strconv.FormatFloat(self.endB.ComparableValue(), 'f', -1, 64)
	}

/**
 * Range of the form a < n < b 
 */
func RangeClosedExclusive(a, b ComparableR) Range {
	return Range{inclusiveA: false, inclusiveB: false, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a <= n <= b 
 */
func RangeClosedInclusive(a, b ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a <= n < b 
 */
func RangeClosedInclusiveLeft(a, b ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: false, openA: false, openB: false, endA: a, endB: b}
}

/**
 * Range of the form a < n <= b 
 */
func RangeClosedInclusiveRight(a, b ComparableR) Range {
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
func RangeOpenLeftInclusive(a ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: true, endA: a, endB: nil}
}

/**
 * Range of the form a < n
 */
func RangeOpenLeftExclusive(a ComparableR) Range {
	return Range{inclusiveA: false, inclusiveB: true, openA: false, openB: true, endA: a, endB: nil}
}

/**
 * Range of the form n <= b 
 */
func RangeOpenRightInclusive(a ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: true, openB: false, endA: nil, endB: a}
}

/**
 * Range of the form n < b
 */
func RangeOpenRightExclusive(a ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: false, openA: true, openB: false, endA: nil, endB: a}
}

/**
 * Range Containing a single value
 */
func RangeSingleton(a ComparableR) Range {
	return Range{inclusiveA: true, inclusiveB: true, openA: false, openB: false, endA: a, endB: a}
}

