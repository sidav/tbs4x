package main

type player struct {
	exploredTiles [][]bool
}

func (p *player) init(mapW, mapH int) {
	p.exploredTiles = make([][]bool, mapW)
	for i := range p.exploredTiles {
		p.exploredTiles[i] = make([]bool, mapH)
	}
}

func (p *player) exploreAround(s *scene, x, y, dist int) {
	for i := x - dist; i <= x+dist; i++ {
		for j := y - dist; j <= y+dist; j++ {
			if s.areCoordsValid(i, j) {
				p.exploredTiles[i][j] = true
			}
		}
	}
}
