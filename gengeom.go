package gengeom

import (
	math "github.com/gabe-lee/genmath"
	vecs "github.com/gabe-lee/genvecs"
)

func PointOnCircle[T_ROT math.Real, T_VEC math.Real](radians T_ROT, radius T_VEC, center vecs.Vec2[T_VEC]) vecs.Vec2[T_VEC] {
	sin := math.Sin(float64(radians))
	cos := math.Cos(float64(radians))
	x, y := T_VEC(cos*float64(radius)), T_VEC(sin*float64(radius))
	return vecs.Vec2[T_VEC]{x, y}.Add(center)
}
func PointOnCircleDeg[T_ROT math.Real, T_VEC math.Real](degrees T_ROT, radius T_VEC, center vecs.Vec2[T_VEC]) vecs.Vec2[T_VEC] {
	return PointOnCircle(float64(degrees)*math.RAD2DEG, radius, center)
}

func PointsOnCircle[T_ROT math.Real, T_VEC math.Real, I_CNT math.Integer](rotationRadians T_ROT, radius T_VEC, center vecs.Vec2[T_VEC], count I_CNT) []vecs.Vec2[T_VEC] {
	count = math.Floor(count)
	thetaStep := math.TAU / float64(count)
	points := make([]vecs.Vec2[T_VEC], count)
	for i, t := I_CNT(0), float64(rotationRadians); i < count; i, t = i+1, t+thetaStep {
		points[i] = PointOnCircle(t, radius, center)
	}
	return points
}
func PointsOnCircleDeg[T_ROT math.Real, T_VEC math.Real, I_CNT math.Integer](degrees T_ROT, radius T_VEC, center vecs.Vec2[T_VEC], count I_CNT) []vecs.Vec2[T_VEC] {
	return PointsOnCircle(float64(degrees)*math.RAD2DEG, radius, center, count)
}

func PointsOnRing[T_ROT math.Real, T_VEC math.Real, I_CNT math.Integer](rotationRadians T_ROT, radiusInner T_VEC, radiusOuter T_VEC, center vecs.Vec2[T_VEC], count I_CNT) []vecs.Vec2[T_VEC] {
	count = math.Floor(count)
	thetaStep := math.TAU / float64(count)
	points := make([]vecs.Vec2[T_VEC], count*2)
	for i, t := I_CNT(0), float64(rotationRadians); i < count*2; i, t = i+2, t+thetaStep {
		points[i] = PointOnCircle(t, radiusInner, center)
		points[i+1] = PointOnCircle(t, radiusOuter, center)
	}
	return points
}
func PointsOnRingDeg[T_ROT math.Real, T_VEC math.Real, I_CNT math.Integer](degrees T_ROT, radiusInner T_VEC, radiusOuter T_VEC, center vecs.Vec2[T_VEC], count I_CNT) []vecs.Vec2[T_VEC] {
	return PointsOnRing(float64(degrees)*math.RAD2DEG, radiusInner, radiusOuter, center, count)
}

func Circumference[T math.Real](radius T) T {
	return T(math.TAU * float64(radius))
}

func AreaOfTriangle[T math.Real](v vecs.Tri2[T]) T {
	return math.Abs(((v[1].X()*v[0].Y())-(v[0].X()*v[1].Y()))+((v[2].X()*v[1].Y())-(v[1].X()*v[2].Y()))+((v[0].X()*v[2].Y())-(v[2].X()*v[0].Y()))) / 2
}

func AreaOfCircle[T math.Real](radius T) T {
	return T(math.PI * math.Pow(float64(radius), 2))
}
