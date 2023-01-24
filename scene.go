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
