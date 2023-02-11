package main

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
			// perform all production
			for _, c := range s.cities {
				s.performProductionForCity(c)
			}
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
			s.currentTurn++
		}

	}
}
