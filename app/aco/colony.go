package aco

type Colony struct {
	Params Parameter
	Graph  Graph
	Ants   []*Ant
}

type AntRoute struct {
	Length int
	Route  []int
}

func NewColony(params Parameter, graph Graph) *Colony {
	ants := make([]*Ant, params.NumAnts)
	for i := 0; i < params.NumAnts; i++ {
		ants[i] = NewAnt(params, graph)
	}

	colony := &Colony{}
	colony.Params = params
	colony.Graph = graph
	colony.Ants = ants

	return colony
}

func (c *Colony) updateColony() {
	c.constructAnts()
}

func (c *Colony) constructAnts() {
	for _, ant := range c.Ants {
		ant.ConstructRoutes()
	}
}

func (c *Colony) calcPassedPheromone() {
	for _, ant := range c.Ants {
		ant.calcPassedPheromone()
	}
}

func (c *Colony) resetColony() {
	for _, ant := range c.Ants {
		ant.resetAnt()
	}
}

func (c *Colony) GetBestRoutes() AntRoute {
	length := 1e20
	var route []int
	for _, ant := range c.Ants {
		if l := ant.calcLengthRoutes(); l < length {
			length = l
			route = append([]int{}, ant.VisitedRoutes...)
		}
	}

	antRoute := &AntRoute{}
	antRoute.Length = int(length)
	antRoute.Route = route

	return *antRoute
}
