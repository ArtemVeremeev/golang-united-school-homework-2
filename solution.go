package square

import "math"

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

type intCustomType int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pow(sideLen, 2) * math.Pi
	default:
		return 0
	}
}
