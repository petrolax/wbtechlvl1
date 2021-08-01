package point

import "math"

func CreatePoint(x, y float64) *Point {
	return &Point{
		x,
		y,
	}
}

type Point struct {
	x, y float64
}

func (pa Point) GetDistance(pb Point) float64 {
	return math.Sqrt(math.Pow((pb.x-pa.x), 2) + math.Pow((pb.y-pa.y), 2))
}
