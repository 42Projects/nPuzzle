package main

//Point represent a point in a 2D plane, x is the abscissa, y the ordinate
type Point struct {
	x, y int
}

type Direction int

const (
	_ Direction = iota
	Down
	Left
	Right
	Up
)

func generateEndState(size int) Matrix {

	m := make(Matrix, size)
	for k := 0; k < size; k++ {
		m[k] = make([]int, size)
	}

	x, y, direction := 0, 0, Right
	for k := 1; k < size*size; k++ {
		m[y][x] = k
		switch direction {
		case Down:
			if y == size-1 || m[y+1][x] != 0 {
				direction = Left
				x -= 1
			} else {
				y += 1
			}
		case Left:
			if x == 0 || m[y][x-1] != 0 {
				direction = Up
				y -= 1
			} else {
				x -= 1
			}
		case Right:
			if x == size-1 || m[y][x+1] != 0 {
				direction = Down
				y += 1
			} else {
				x += 1
			}
		case Up:
			if y == 0 || m[y-1][x] != 0 {
				direction = Right
				x += 1
			} else {
				y -= 1
			}
		}
	}

	return m
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
