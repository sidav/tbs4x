package main

type scene struct {
	cities []*city
	tiles  [][]*tile
}

func (s *scene) areCoordsValid(x, y int) bool {
	return x > 0 && x < len(s.tiles) && y > 0 && y < len(s.tiles[x])
}

func (s *scene) addCity(c *city) {
	s.cities = append(s.cities, c)
}

func (s *scene) getCityAt(x, y int) *city {
	for _, c := range s.cities {
		if c.x == x && c.y == y {
			return c
		}
	}
	return nil
}

func (s *scene) isTileApplicableForCity(x, y int) bool {
	const rang = 2
	for i := x - rang; i <= x+rang; i++ {
		for j := y - rang; j <= y+rang; j++ {
			if !s.areCoordsValid(i, j) {
				return false
			}
			if s.tiles[i][j].code != TILE_GRASS {
				return false
			}
		}
	}
	return true
}
