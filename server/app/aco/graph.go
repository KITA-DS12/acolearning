package aco

import (
	"math"
)

type Graph struct {
	Params             Parameter
	Coordinates        []Point
	Distances          [][]float64
	TrailPheromones    [][]float64
	NewTrailPheromones [][]float64
	HeuristicValues    [][]float64
}

type Point struct {
	X float64
	Y float64
}

func NewGraph(params Parameter, points []Point) *Graph {

	graph := &Graph{}
	graph.Params = params
	graph.Coordinates = make([]Point, graph.Params.NumNodes)
	graph.Distances = make([][]float64, graph.Params.NumNodes)
	graph.TrailPheromones = make([][]float64, graph.Params.NumNodes)
	graph.NewTrailPheromones = make([][]float64, graph.Params.NumNodes)
	graph.HeuristicValues = make([][]float64, graph.Params.NumNodes)

	for i := 0; i < graph.Params.NumNodes; i++ {
		graph.Distances[i] = make([]float64, graph.Params.NumNodes)
		graph.TrailPheromones[i] = make([]float64, graph.Params.NumNodes)
		graph.NewTrailPheromones[i] = make([]float64, graph.Params.NumNodes)
		graph.HeuristicValues[i] = make([]float64, graph.Params.NumNodes)
	}

	graph.prepareGraph(points)

	return graph
}

func (g *Graph) prepareGraph(points []Point) {
	g.Coordinates = points
	for i := 0; i < g.Params.NumNodes; i++ {
		for j := 0; j < g.Params.NumNodes; j++ {
			g.Distances[i][j] = calcDistance(
				g.Coordinates[i], g.Coordinates[j],
			)
		}
	}

	for i := 0; i < g.Params.NumNodes; i++ {
		for j := 0; j < g.Params.NumNodes; j++ {
			if i < j {
				distance := g.Distances[i][j]
				g.TrailPheromones[i][j] = float64(g.Params.Q) / distance
				g.TrailPheromones[j][i] = float64(g.Params.Q) / distance
				g.HeuristicValues[i][j] = float64(g.Params.Q) / distance
				g.HeuristicValues[j][i] = float64(g.Params.Q) / distance
			}
		}
	}
}

func (g *Graph) resetGraph() {
	g.NewTrailPheromones = make([][]float64, g.Params.NumNodes)
	for i := range g.NewTrailPheromones {
		g.NewTrailPheromones[i] = make([]float64, g.Params.NumNodes)
	}
}

func (g *Graph) resetPheromones() {
	for i := 0; i < g.Params.NumNodes; i++ {
		for j := 0; j < g.Params.NumNodes; j++ {
			if i < j {
				distance := g.Distances[i][j]
				g.TrailPheromones[i][j] = float64(g.Params.Q) / distance
				g.TrailPheromones[j][i] = float64(g.Params.Q) / distance
			}
		}
	}
}

func calcDistance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}
