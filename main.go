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
