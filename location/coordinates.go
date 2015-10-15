package location

import (
	m "math"
	st "strconv"
)

// Immutable construct. 
// 
type Vec3 struct {
	xCoord float64
	yCoord float64
	zCoord float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (self Vec3) GetX() float64 {
	return self.xCoord;
}

func (self Vec3) GetY() float64 {
	return self.yCoord;
}

func (self Vec3) GetZ() float64 {
	return self.zCoord;
}

func (self Vec3) DotProduct(other Vec3) float64 {
	return self.xCoord * other.xCoord + self.yCoord * other.yCoord + self.zCoord * other.zCoord
}

func (self Vec3) Normalize() Vec3 {
	
	sqrt := m.Sqrt(self.xCoord * self.xCoord + self.yCoord * self.yCoord + self.zCoord * self.zCoord)
	
	return NewVec3(self.xCoord / sqrt, self.yCoord / sqrt, self.zCoord / sqrt)
}

func (s Vec3) CrossProduct(o Vec3) Vec3 {
	x := s.yCoord * o.zCoord - o.yCoord * s.zCoord
	y := s.xCoord * o.zCoord - o.xCoord * s.zCoord
	z := s.xCoord * o.yCoord - o.xCoord * s.yCoord
	
	return NewVec3(x, y, z)
}

func (s Vec3) String() string {
	return "(" + st.FormatFloat(s.xCoord, 'f', -1, 64) + "," +
			st.FormatFloat(s.yCoord, 'f', -1, 64) + "," +
			st.FormatFloat(s.zCoord, 'f', -1, 64) + ")"
}


