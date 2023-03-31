package aco

import (
	"math"
	"math/rand"
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

func NewGraph(params Parameter) *Graph {

	g := &Graph{}
	g.Params = params
	g.Coordinates = make([]Point, g.Params.NumNodes)
	g.Distances = make([][]float64, g.Params.NumNodes)
	g.TrailPheromones = make([][]float64, g.Params.NumNodes)
	g.NewTrailPheromones = make([][]float64, g.Params.NumNodes)
	g.HeuristicValues = make([][]float64, g.Params.NumNodes)

	for i := 0; i < g.Params.NumNodes; i++ {
		g.Distances[i] = make([]float64, g.Params.NumNodes)
		g.TrailPheromones[i] = make([]float64, g.Params.NumNodes)
		g.NewTrailPheromones[i] = make([]float64, g.Params.NumNodes)
		g.HeuristicValues[i] = make([]float64, g.Params.NumNodes)
	}

	g.prepareGraph()

	return g
}

func (g *Graph) prepareGraph() {
	for i := 0; i < g.Params.NumNodes; i++ {
		g.Coordinates[i] = Point{
			X: rand.Float64() * 1000,
			Y: rand.Float64() * 1000,
		}
	}
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
