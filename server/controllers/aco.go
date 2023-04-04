package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/KITA-DS12/acolearning.git/app/aco"
	"github.com/KITA-DS12/acolearning.git/middleware"
)

type Graph struct {
	Nodes   map[string]Node `json:"nodes"`
	Edges   map[string]Edge `json:"edges"`
	Layouts Layout          `json:"layouts"`
}

type Node struct {
	Name string `json:"name"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Layout struct {
	Nodes map[string]Point `json:"nodes"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func runAco(w http.ResponseWriter, r *http.Request) {
	middleware.CorsMiddleware(w)
	if r.Method == "OPTIONS" {
		middleware.OptionMiddleware(w)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	var data []byte

	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var graph Graph
	if err := json.Unmarshal(data, &graph); err != nil {
		fmt.Println(err)
		return
	}

	var layouts []aco.Point
	for i := 0; i < len(graph.Layouts.Nodes); i++ {
		var p aco.Point
		key := "node" + strconv.Itoa(i)
		p.X = graph.Layouts.Nodes[key].X
		p.Y = graph.Layouts.Nodes[key].Y
		layouts = append(layouts, p)
	}

	params := aco.Parameter{
		NumAnts:       30,
		NumNodes:      20,
		Q:             100,
		Alpha:         3,
		Beta:          5,
		Rou:           0.9,
		MaxIteration:  500,
		MinTau:        0,
		MaxTau:        3000,
		NumCycleReset: 30,
		IndexStart:    0,
	}

	solver := aco.NewSolver(params, layouts)
	route := solver.RunAco()
	var edges = make(map[string]Edge)
	for i := 0; i < len(route)-1; i++ {
		var name = "edge" + strconv.Itoa(i)
		var edge Edge
		edge.Source = "node" + strconv.Itoa(route[i])
		edge.Target = "node" + strconv.Itoa(route[i+1])
		edges[name] = edge
	}

	graph.Edges = edges

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(graph)
}
