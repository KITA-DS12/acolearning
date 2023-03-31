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

	return g
}
