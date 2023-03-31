package aco

type Parameter struct {
	NumAnts       int
	NumNodes      int
	Q             int
	Alpha         float64
	Beta          float64
	Rou           float64
	MaxIteration  int
	MinTau        int
	MaxTau        int
	NumCycleReset int
	IndexStart    int
}
