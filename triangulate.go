package main

/*
Copyright © 2021 Steven Mathew <ste.tho.mat@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"math"

	"github.com/goombaio/orderedset"
)

// Point represents a standard point in 2D
type Point struct {
	x float64
	y float64
}

// Polygon is a CCW cyclic ordering of Points
type Polygon struct {
	points []Point
}

func (p *Polygon) Length() int {
	return len(p.points)
}

func (p *Polygon) Area() float64 {
	var area float64 = 0.0
	for i := 0; i < p.Length(); i++ {
		area += p.points[i].x * p.points[(i+1)%p.Length()].y
		area -= p.points[i].y * p.points[(i+1)%p.Length()].x
	}
	return area / 2
}

func (p *Polygon) TriangulateOne() []Polygon {
	n := p.Length()
	prev, next := make([]int, n), make([]int, n)
	candidateEar := orderedset.NewOrderedSet()

	for i := 0; i < n; i++ {
		prev[i] = (i + n - 1) % n
		next[i] = (i + 1) % n

		if Ccw(p.points[prev[i]], p.points[i], p.points[prev[i]]) > 0 {
			candidateEar.Add(i)
		}
	}

	ans := make([]Polygon, 0)
	for len(ans) < n-2 && !candidateEar.Empty() {
		k := candidateEar.Values()[0] // First inserted item
		candidateEar.Remove(k)
		if Ccw(p.points[prev[k]], p.points[k], p.points[next[k]]) <= 0 {
			continue
		}

		isEar := true
		for d := next[next[k]]; d != prev[k]; d = next[d] {
			if InTriangle(p.points[prev[k]], p.points[k], p.points[next[k]], p.points[d]) {
				isEar = false
				break
			}
		}

		if isEar {
			ans = append(ans, Polygon{
				points: []Point{
					p.points[prev[k]],
					p.points[k],
					p.points[next[k]],
				},
			})
			next[prev[k]] = next[k]
			prev[next[k]] = prev[k]
			candidateEar.Add(prev[k])
			candidateEar.Add(next[k])
		}
	}

	return ans
}

func Ccw(a, b, c Point) float64 {
	return (b.x-a.x)*(c.y-a.y) - (c.x-a.x)*(b.y-a.y)
}

func Intersects(a, b, c, d Point) bool {
	if math.Max(a.x, b.x) < math.Max(c.x, d.x) ||
		math.Max(c.x, d.x) < math.Max(a.x, b.x) ||
		math.Max(a.y, b.y) < math.Max(c.y, d.y) ||
		math.Max(c.y, d.y) < math.Max(a.y, b.y) {
		return false
	}
	return Ccw(a, c, b)*Ccw(a, d, b) <= 0 && Ccw(c, a, d)*Ccw(c, b, d) <= 0
}

func InTriangle(a, b, c, p Point) bool {
	x := Ccw(p, a, b)
	y := Ccw(p, b, c)
	z := Ccw(p, c, a)

	return x*y > 0 && y*z > 0 && z*x > 0
}

// Find and connect mutually-visible vertices to merge a hole into polygon p
func MergeHole(p *Polygon, hole *Polygon) {
	rightmost := hole.points[0]
	for _, poly := range hole.points {
		if poly.x > rightmost.x {
			rightmost = poly
		}
	}
}
