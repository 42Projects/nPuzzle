package npuzzle

func getMove(before, after Matrix) byte {

	zeroTileBefore := before.getTilePosition(0)
	zeroTileAfter := after.getTilePosition(0)
	switch {
	case zeroTileAfter.x-zeroTileBefore.x == 1:
		return 'R'
	case zeroTileBefore.x-zeroTileAfter.x == 1:
		return 'L'
	case zeroTileAfter.y-zeroTileBefore.y == 1:
		return 'D'
	default:
		return 'U'
	}
}

func mapPath(path []byte, index int, current, parent *Item) {

	path[index-1] = getMove(parent.M, current.M)
	if parent.Parent != nil {
		mapPath(path, index-1, parent, parent.Parent)
	}
}

//StringifyPath maps the solution path to a string, using U (Up), L (Left), R (Right), D (Down)
func StringifyPath(solution *Item) string {

	path := make([]byte, solution.Cost)
	mapPath(path, solution.Cost, solution, solution.Parent)
	return string(path)
}
