package aco

type Graph struct {
	Params             Parameter
	Coordinates        []Point
	Distances          [][]int
	TrailPheromones    [][]int
	NewTrailPheromones [][]int
	HeuristicValues    [][]int
}

type Point struct {
	X int
	Y int
}

func NewGraph(
	params Parameter,
) *Graph {

	g := &Graph{}
	g.Params = params
	g.Coordinates = make([]Point, g.Params.NumNodes)
	g.Distances = make([][]int, g.Params.NumNodes)
	g.TrailPheromones = make([][]int, g.Params.NumNodes)
	g.NewTrailPheromones = make([][]int, g.Params.NumNodes)
	g.HeuristicValues = make([][]int, g.Params.NumNodes)

	for i := 0; i < g.Params.NumNodes; i++ {
		g.Distances[i] = make([]int, g.Params.NumNodes)
		g.TrailPheromones[i] = make([]int, g.Params.NumNodes)
		g.NewTrailPheromones[i] = make([]int, g.Params.NumNodes)
		g.HeuristicValues[i] = make([]int, g.Params.NumNodes)
	}

	return g
}
