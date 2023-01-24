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
	return s.countTilesAllowingBuildingAround(x, y, rang) >= (2*rang+1)*(2*rang+1)
}

func (s *scene) countTilesAllowingBuildingAround(x, y, dist int) int {
	count := 0
	for i := x - dist; i <= x+dist; i++ {
		for j := y - dist; j <= y+dist; j++ {
			if s.areCoordsValid(i, j) && s.tiles[i][j].allowsBuilding() {
				count++
			}
		}
	}
	return count
}
