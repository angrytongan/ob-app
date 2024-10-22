package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Point struct {
	R int `json:"r"`
	X int `json:"x"`
	Y int `json:"y"`
}

func (app *Application) RandomPoints(w http.ResponseWriter, r *http.Request) {
	const maxR = 10
	const maxX = 100
	const maxY = 100

	// Generate some random points.
	numPoints, err := strconv.Atoi(r.PathValue("num"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	out := []Point{}

	for i := 0; i < numPoints; i++ {
		out = append(out, Point{
			R: 1 + rand.Intn(maxR-1), // at least radius of 2 pixels
			X: rand.Intn(maxX),
			Y: rand.Intn(maxY),
		})
	}

	// Send to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)
}
