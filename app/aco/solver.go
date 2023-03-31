package aco

import (
	"fmt"
	"reflect"
)

type Solver struct {
	Params          Parameter
	Graph           Graph
	Colony          Colony
	BestAnt         AntRoute
	GoodAnt         AntRoute
	StagnationCount int
}

func NewSolver(params Parameter) *Solver {
	graph := NewGraph(params)
	colony := NewColony(params, *graph)
	solver := &Solver{}
	solver.Params = params
	solver.Graph = *graph
	solver.Colony = *colony
	solver.BestAnt = AntRoute{}
	solver.GoodAnt = AntRoute{}
	solver.StagnationCount = 0

	return solver
}

func (s *Solver) RunAco() {
	for t := 0; t < s.Params.MaxIteration; t++ {
		s.runCycle()
		s.resetAco()

		if reflect.DeepEqual(s.BestAnt, AntRoute{}) || s.GoodAnt.Length < s.BestAnt.Length {
			s.BestAnt = s.GoodAnt
		} else {
			s.StagnationCount++
		}

		if s.StagnationCount > s.Params.NumCycleReset {
			s.Graph.resetPheromones()
			s.StagnationCount = 0
		}

		fmt.Println(t, "iter, length: ", s.BestAnt.Length)
	}
}

func (s *Solver) runCycle() {
	s.Colony.updateColony()
	s.updatePheromones()
	s.GoodAnt = s.Colony.GetBestRoutes()
}

func (s *Solver) resetAco() {
	s.Colony.resetColony()
	s.Graph.resetGraph()
}

func (s *Solver) updatePheromones() {
	for i := 0; i < s.Params.NumNodes; i++ {
		for j := 0; j < s.Params.NumNodes; j++ {
			s.Graph.TrailPheromones[i][j] =
				s.Graph.TrailPheromones[i][j]*s.Params.Rou + s.Graph.NewTrailPheromones[i][j]
		}
	}
	for i := 0; i < s.Params.NumNodes; i++ {
		for j := 0; j < s.Params.NumNodes; j++ {
			s.Graph.TrailPheromones[i][j] =
				min(float64(s.Params.MaxTau),
					max(s.Graph.TrailPheromones[i][j], float64(s.Colony.Params.MinTau)))
		}
	}
}

func min(v1, v2 float64) float64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 float64) float64 {
	if v1 > v2 {
		return v1
	}
	return v2
}
