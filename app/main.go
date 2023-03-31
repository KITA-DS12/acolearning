package main

import (
	"github.com/KITA-DS12/acolearning.git/aco"
)

func main() {
	params := aco.Parameter{
		NumAnts:       30,
		NumNodes:      300,
		Q:             100,
		Alpha:         3,
		Beta:          5,
		Rou:           0.9,
		MaxIteration:  300,
		MinTau:        0,
		MaxTau:        3000,
		NumCycleReset: 30,
		IndexStart:    0,
	}

	solver := aco.NewSolver(params)
	solver.RunAco()
}
