package main

type Point struct {
	x, y int
}

func (m Matrix) getTilePosition() Point {

	for y, row := range m {
		for x, it := range row {
			if it == 0 {
				return Point{x, y}
			}
		}
	}

	return Point{-1, -1}
}

func (m Matrix) slideDown() Matrix {

	p := m.getTilePosition()
	if p.y == len(m)-1 {
		return nil
	}

	m[p.y][p.x], m[p.y+1][p.x] = m[p.y+1][p.x], m[p.y][p.x]
	return m
}

func (m Matrix) slideLeft() Matrix {

	p := m.getTilePosition()
	if p.x == 0 {
		return nil
	}

	m[p.y][p.x], m[p.y][p.x-1] = m[p.y][p.x-1], m[p.y][p.x]
	return m
}

func (m Matrix) slideRight() Matrix {

	p := m.getTilePosition()
	if p.x == len(m[0])-1 {
		return nil
	}

	m[p.y][p.x], m[p.y][p.x+1] = m[p.y][p.x+1], m[p.y][p.x]
	return m
}

func (m Matrix) slideUp() Matrix {

	p := m.getTilePosition()
	if p.y == 0 {
		return nil
	}

	m[p.y][p.x], m[p.y-1][p.x] = m[p.y-1][p.x], m[p.y][p.x]
	return m
}

func (m Matrix) solve() Matrix {

	return m
}
