package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func CountDistance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))
}

func task24() {
	a := NewPoint(2, -5)
	b := NewPoint(-4, 3)
	fmt.Println(CountDistance(a, b))
}
