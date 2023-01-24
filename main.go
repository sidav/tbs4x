package main

import (
	"tbs4x/lib/random"
	"tbs4x/lib/tcell_console_wrapper"
)

var cw tcell_console_wrapper.ConsoleWrapper
var rnd random.PRNG

var GAME_RUNS = true

func main() {
	//rnd = pcgrandom.New(-1)
	//perlinNoiseMap := perlin_noise.GeneratePerlinNoiseMap(1280, 800, 8, 0.005, rnd)
	//perlin_noise.MakeNoiseRectangular(perlinNoiseMap)
	//perlin_noise.CreatePerlinNoiseImage(perlinNoiseMap)
	//return

	cw.Init()
	defer cw.Close()

	s := &scene{}
	s.init()
	rend := newAsciiRenderer()
	pc := &playerController{}

	for GAME_RUNS {
		rend.renderMainScreen(s, pc)
		pc.playerControl()
	}
}
