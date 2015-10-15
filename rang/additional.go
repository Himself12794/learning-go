package rang

type Float64 float64
func (s Float64) ComparableValue() float64 {
	return float64(s)
}


type CompBase struct {
	v float64
}
// Method implementation
	func (self CompBase) ComparableValue() float64 {
		return self.v;
	}

// Static methods - these allow translation of native numbers to be comparable
func Byte(v byte) ComparableR {
	return &CompBase{float64(v)}
}

func Int16(v int16) ComparableR {
	return &CompBase{float64(v)}
}

func Int(v int) ComparableR {
	return &CompBase{float64(v)}
}

func Float32(v float32) ComparableR {
	return &CompBase{float64(v)}
}

/*func Float64(v float64) ComparableR {
	return &CompBase{v}
}*/

func UInt(v uint) ComparableR {
	return &CompBase{float64(v)}
}