package state

import (
	"sync"

	"smashrift.mikekaipis.com/utils"
)

func Init() {
	players = sync.Map{}
	startGameLoop()
}

func AddPlayer(player Player) error {
	if !player.Validate() {
		return ErrorInvalidPlayer
	}
	players.Store(player.ID, player)
	return nil
}

func GetPlayer(id string) (Player, error) {
	player, ok := players.Load(id)
	if !ok {
		return Player{}, ErrorPlayerNotFound
	}
	return player.(Player), nil
}

func GetPlayers() []Player {
	totalPlayers := []Player{}
	players.Range(func(key, value interface{}) bool {
		totalPlayers = append(totalPlayers, value.(Player))
		return true
	})
	return totalPlayers
}

func UpdatePlayer(ID string, direction utils.Vector2, speed *float64) error {
	player, ok := GetPlayer(ID)
	if ok != nil {
		return ErrorPlayerNotFound
	}
	player.Direction = direction.Normalize()
	if speed != nil {
		player.Speed = *speed
	}
	players.Store(ID, player)
	return nil
}

func DeletePlayer(ID string) error {
	players.Delete(ID)
	return nil
}
