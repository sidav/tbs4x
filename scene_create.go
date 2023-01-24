package main

func (s *scene) init() {
	testMap := []string{
		"~~~~~~~~~~~~~~~",
		"~.............~",
		"~...^^........~",
		"~.......~~....~",
		"~.............~",
		"~.............~",
		"~.............~",
		"~.............~",
		"~.............~",
		"~~~~~~~~~~~~~~~",
	}
	s.tiles = make([][]*tile, len(testMap))
	for i := range s.tiles {
		s.tiles[i] = make([]*tile, len(testMap[i]))
	}
	for x := range testMap {
		for y := range testMap[x] {
			switch rune(testMap[x][y]) {
			case '~':
				s.tiles[x][y] = &tile{code: TILE_WATER}
			case '.':
				s.tiles[x][y] = &tile{code: TILE_SAND}
			case '^':
				s.tiles[x][y] = &tile{code: TILE_ROCK}
			}
		}
	}
}
