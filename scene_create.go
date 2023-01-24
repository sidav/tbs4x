package main

import "tbs4x/lib/perlin_noise"

func (s *scene) init() {
	// Perlin noise generation
	noiseMap := perlin_noise.GeneratePerlinNoiseMap(64, 64, 20, 8, rnd)
	perlin_noise.MakeNoiseRectangular(noiseMap)
	s.tiles = make([][]*tile, len(noiseMap))
	for i := range s.tiles {
		s.tiles[i] = make([]*tile, len(noiseMap[i]))
	}
	for x := range noiseMap {
		for y := range noiseMap[x] {
			switch {
			case noiseMap[x][y] < 0.5:
				s.tiles[x][y] = &tile{code: TILE_WATER}
			case noiseMap[x][y] < 0.55:
				s.tiles[x][y] = &tile{code: TILE_SAND}
			case noiseMap[x][y] < 0.90:
				s.tiles[x][y] = &tile{code: TILE_GRASS}
			default:
				s.tiles[x][y] = &tile{code: TILE_MOUNTAIN}
			}
		}
	}
}
