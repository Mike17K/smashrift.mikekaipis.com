package state

import (
	"math/rand"
	"sync"

	"smashrift.mikekaipis.com/utils"
)

type Player struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Position  utils.Vector2 `json:"position"`
	Direction utils.Vector2 `json:"direction"`
	Speed     float64       `json:"speed"`
}

func (p Player) Validate() bool {
	return p.Direction.Magnitude() <= 1.001 && p.Speed >= 0 && p.Name != "" && p.ID != ""
}

func CreateDummyPlayer() Player {
	return Player{
		ID:       utils.RandSeq(10),
		Name:     "dummy",
		Position: utils.Vector2{X: 2500, Y: 2500},
		Direction: utils.Vector2{
			X: rand.Float64()*2 - 1,
			Y: rand.Float64()*2 - 1,
		}.Normalize(),
		Speed: 2,
	}
}

type Map struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

var players = sync.Map{}
var mapData = Map{
	Width:  5000,
	Height: 5000,
}

func GetRandMapPosition() utils.Vector2 {
	return utils.Vector2{
		X: rand.Float64() * mapData.Width,
		Y: rand.Float64() * mapData.Height,
	}
}
