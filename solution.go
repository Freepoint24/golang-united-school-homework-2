package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNum2 int

const SidesTriangle sidesNum2 = 3
const SidesSquare sidesNum2 = 4
const SidesCircle sidesNum2 = 0

func CalcSquare(sideLen float64, sidesNum sidesNum2) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}

}
