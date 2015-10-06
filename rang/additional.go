package rang

import (

)

type CompBase struct {
	v float64
}
// Method implementation
	func (self CompBase) Value() float64 {
		return self.v;
	}

// Static methods
func Byte(v byte) Comparable {
	return &CompBase{float64(v)}
}

func Int16(v int16) Comparable {
	return &CompBase{float64(v)}
}

func Int(v int) Comparable {
	return &CompBase{float64(v)}
}

func Float32(v float32) Comparable {
	return &CompBase{float64(v)}
}

func Float64(v float64) Comparable {
	return &CompBase{v}
}

func UInt(v uint) Comparable {
	return &CompBase{float64(v)}
}
