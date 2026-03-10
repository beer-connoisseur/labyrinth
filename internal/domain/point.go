package domain

func (p *Point) GetNeighbors(directions [][2]int) []Point {
	neighbors := make([]Point, 0)
	for _, direction := range directions {
		neighbor := Point{X: p.X + direction[0], Y: p.Y + direction[1]}
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}
