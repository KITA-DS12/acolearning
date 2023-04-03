package aco

import (
	"math"
	"math/rand"
)

type Ant struct {
	Params        Parameter
	Graph         Graph
	VisitedRoutes []int
	IsVisited     []bool
}

func NewAnt(params Parameter, graph Graph) *Ant {
	ant := &Ant{}
	ant.Params = params
	ant.Graph = graph
	ant.VisitedRoutes = []int{ant.Params.IndexStart}
	ant.IsVisited = make([]bool, ant.Params.NumNodes)
	ant.IsVisited[ant.Params.IndexStart] = true

	return ant
}

func (ant *Ant) ConstructRoutes() {
	ant.VisitedRoutes = append(ant.VisitedRoutes, ant.Params.IndexStart)
	ant.IsVisited[ant.Params.IndexStart] = true
	for i := 0; i < ant.Params.NumNodes-1; i++ {
		v := ant.VisitedRoutes[len(ant.VisitedRoutes)-1]
		toNodes, toProbabilities := ant.calcProbabilityFromV(v)

		var toNode int
		if rand.Float64() < 0.1 {
			toNode = toNodes[rand.Intn(len(toNodes))]
		} else {
			randomProbability := rand.Float64()
			toNode = toNodes[getInsertionIndex(toProbabilities, randomProbability)]
		}
		ant.VisitedRoutes = append(ant.VisitedRoutes, toNode)
		ant.IsVisited[toNode] = true
	}
}

func (ant *Ant) calcProbabilityFromV(v int) ([]int, []float64) {
	sumProbability := 0.0

	toNodes := []int{}
	toPheromones := []float64{}

	for toNode := 0; toNode < ant.Params.NumNodes; toNode++ {
		if toNode == v || ant.IsVisited[toNode] {
			continue
		}
		pheromone := math.Pow(ant.Graph.TrailPheromones[v][toNode], ant.Params.Alpha) * math.Pow(ant.Graph.HeuristicValues[v][toNode], ant.Params.Beta)
		sumProbability += pheromone
		toNodes = append(toNodes, toNode)
		toPheromones = append(toPheromones, pheromone)
	}

	toProbabilities := make([]float64, len(toPheromones))
	for i, p := range toPheromones {
		toProbabilities[i] = p / sumProbability
	}

	for i := range toProbabilities[1:] {
		toProbabilities[i+1] += toProbabilities[i]
	}

	return toNodes, toProbabilities
}

func getInsertionIndex(arr []float64, value float64) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if value == arr[mid] {
			return mid
		} else if value < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func (ant *Ant) calcLengthRoutes() float64 {
	lengthRoute := 0.0
	for i := 0; i < ant.Params.NumNodes-1; i++ {
		lengthRoute += ant.Graph.Distances[ant.VisitedRoutes[i]][ant.VisitedRoutes[i+1]]
	}

	return lengthRoute
}

func (ant *Ant) calcPassedPheromone() {
	lengthRoute := ant.calcLengthRoutes()
	q := float64(ant.Params.Q)
	for i := 0; i < ant.Params.NumNodes-1; i++ {
		ant.Graph.NewTrailPheromones[ant.VisitedRoutes[i]][ant.VisitedRoutes[i+1]] += q / lengthRoute
	}
}

func (ant *Ant) resetAnt() {
	ant.IsVisited = make([]bool, ant.Params.NumNodes)
	ant.VisitedRoutes = []int{}
}
