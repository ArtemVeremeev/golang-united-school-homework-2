package square

import "math"

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

type intCustomType int

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
