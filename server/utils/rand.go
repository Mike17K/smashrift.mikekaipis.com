package utils

import (
	"math"
	"math/rand"
)

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vector2) Multiply(scalar float64) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Normalize() Vector2 {
	magnitude := v.Magnitude()
	return Vector2{X: v.X / magnitude, Y: v.Y / magnitude}
}

func GetRandDirection() Vector2 {
	return Vector2{
		X: rand.Float64()*2 - 1,
		Y: rand.Float64()*2 - 1,
	}.Normalize()
}
