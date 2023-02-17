package main

type player struct {
	endedThisTurn         bool
	exploredTiles         [][]bool
	seenTiles             [][]bool
	notificationsThisTurn []string

	currentMoney int
}

func (p *player) init(mapW, mapH int) {
	p.exploredTiles = make([][]bool, mapW)
	for i := range p.exploredTiles {
		p.exploredTiles[i] = make([]bool, mapH)
	}
	p.seenTiles = make([][]bool, mapW)
	for i := range p.seenTiles {
		p.seenTiles[i] = make([]bool, mapH)
	}
}

func (p *player) hasNotifications() bool {
	return len(p.notificationsThisTurn) > 0
}

func (p *player) clearNotifications() {
	p.notificationsThisTurn = p.notificationsThisTurn[:0]
}

func (p *player) addNotification(n string) {
	p.notificationsThisTurn = append(p.notificationsThisTurn, n)
}

func (p *player) resetVision() {
	for i := range p.seenTiles {
		for j := range p.seenTiles[i] {
			p.seenTiles[i][j] = false
		}
	}
}

func (p *player) exploreAround(s *scene, x, y, dist int) {
	for i := x - dist; i <= x+dist; i++ {
		for j := y - dist; j <= y+dist; j++ {
			if s.areCoordsValid(i, j) {
				p.exploredTiles[i][j] = true
				p.seenTiles[i][j] = true
			}
		}
	}
}
