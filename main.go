package main

import (
	"tbs4x/lib/console/tcell_console_wrapper"
	"tbs4x/lib/random"
	"tbs4x/lib/random/pcgrandom"
)

var cw tcell_console_wrapper.ConsoleWrapper
var rnd random.PRNG

var GAME_RUNS = true
var rend *asciiRenderer

func main() {
	rnd = pcgrandom.New(-1)
	//perlinNoiseMap := perlin_noise.GeneratePerlinNoiseMap(1280, 800, 64, 8, rnd)
	//perlin_noise.MakeNoiseRectangular(perlinNoiseMap)
	//perlin_noise.CreatePerlinNoiseImage(perlinNoiseMap)
	//return

	cw.Init()
	defer func() {
		cw.Close()
		if x := recover(); x != nil {
			panic(x)
		}
	}()

	gameLoop()
}

func gameLoop() {
	s := &scene{}
	s.init(32, 32)
	s.performExploration()
	rend = newAsciiRenderer()
	pc := &playerController{controlsPlayer: s.players[0]}
	pc.resetMode()
	pc.setCursorAt(s.cities[0].x, s.cities[0].y)
	for GAME_RUNS {
		for _, currPlayer := range s.players {
			// beginning of turn cleanup
			currPlayer.endedThisTurn = false
			for i := range s.units {
				if s.units[i].owner == currPlayer {
					s.units[i].movementPointsRemaining = s.units[i].getStaticData().geoscapeStats.speed
				}
			}

			for !currPlayer.endedThisTurn && GAME_RUNS {
				rend.renderMainScreen(s, pc)
				pc.playerControl(s)
			}
		}

	}
}
