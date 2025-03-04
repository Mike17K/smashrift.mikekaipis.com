package state

import (
	"time"

	"smashrift.mikekaipis.com/benchmarking"
)

func updatePlayers() {
	// Update all players
	players.Range(func(key, value interface{}) bool {
		player := value.(Player)

		player.Position = player.Position.Add(player.Direction.Multiply(player.Speed))

		// Check if the player is out of bounds
		if player.Position.X < 0 {
			player.Position.X = 0
		}
		if player.Position.X > mapData.Width {
			player.Position.X = mapData.Width
		}
		if player.Position.Y < 0 {
			player.Position.Y = 0
		}
		if player.Position.Y > mapData.Height {
			player.Position.Y = mapData.Height
		}

		players.Store(key, player)
		return true
	})
}

func startGameLoop() {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 33) // Set the ticker for 30 FPS
		defer ticker.Stop()                             // Ensure the ticker stops when the goroutine exits

		loops := 0
		timerStart := time.Now()
		for range ticker.C {
			loops++
			if loops%100 == 0 {
				benchmarking.SetMetric("game_loop", float64(loops))
				benchmarking.SetMetric("fps", float64(loops)/time.Since(timerStart).Seconds())
			}
			updatePlayers()
		}
		benchmarking.SetMetric("game_loop", float64(loops))
		benchmarking.SetMetric("fps", float64(loops)/time.Since(timerStart).Seconds())
	}()
}
