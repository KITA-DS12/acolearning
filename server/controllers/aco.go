package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	X int `json:"x"`
	Y int `json:"y"`
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

	log.Println("--- Graph ---")
	log.Println("--- Nodes ---")
	for id, n := range graph.Nodes {
		log.Printf("%s: %s\n", id, n.Name)
	}
	log.Println("--- Edges ---")
	for id, e := range graph.Edges {
		log.Printf("%s: %s -> %s\n", id, e.Source, e.Target)
	}
	log.Println("--- Layouts ---")
	for i, p := range graph.Layouts.Nodes {
		log.Printf("%s: {X:%d, Y:%d}\n", i, p.X, p.Y)
	}
}
