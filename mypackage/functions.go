package mypackage

import "math"

func distance(p Point, q Point) float64 {
	return math.Hypot(float64(p.X-p.Y), float64(q.X-q.Y))
}
