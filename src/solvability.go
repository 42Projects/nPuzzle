package main

/*
 * Because of the snail goal pattern, our algorithm to find solvability is a bit different than usual.
 * We invert the standard solvability for size=3 to size=6, then we revert back to normal for the next 4 sizes, and so
 * on... A size which standard solvability is inverted is defined as an odd size
 */
func isSolvable(m Matrix) bool {

	size := len(m)

	/* Look for the size oddity of our current size, if so we would need to invert the result of the algorithm */
	sizeOddity := (size-2)/4%2 == 0

	line := make([]int, size*size-1)
	k := 0
	for _, row := range m {
		for _, num := range row {
			if num != 0 {
				line[k] = num
				k++
			}
		}
	}

	inversion := 0
	for k, num1 := range line {
		for _, num2 := range line[k+1:] {
			if num1 > num2 {
				inversion++
			}
		}
	}

	/* If grid size is odd, the puzzle is only solvable if the number of inversion is odd as well */
	if size%2 != 0 {
		return (inversion%2 != 0) == sizeOddity
	}

	/* Otherwise, we also need to check the blank tile position. */
	blankTile := m.getTilePosition(0)

	return ((inversion%2 != 0 && blankTile.y%2 == 0) || (inversion%2 == 0 && blankTile.y%2 != 0)) == sizeOddity
}
