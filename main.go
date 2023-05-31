package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/FedoseevAlex/bandits/pb"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/FedoseevAlex/bandits/internal/bandits"
	"github.com/FedoseevAlex/bandits/internal/ucb"
)

type BanditData struct {
	TotalRounds int
	Candidates  map[string]float64
}

var banditData = bandits.ContextualData{
	Rounds: 0,
	Rewards: map[bandits.ActionID]float64{
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
	},
}
var strategy bandits.Strateger

func main() {
	http.Handle("/", http.HandlerFunc(RootHandler))
	http.Handle("/show", http.HandlerFunc(ShowHandler))
	http.Handle("/click", http.HandlerFunc(ClickHandler))

	strategy = ucb.NewUCB1Strategy()

	log.Println("Starting server on port 8080")
	http.ListenAndServe("localhost:8080", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bodyBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ShowHandler(w http.ResponseWriter, req *http.Request) {
	candidateToShow, err := strategy.Choose(context.Background(), &banditData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("data: %+v", banditData)
	w.Write([]byte(candidateToShow))
}

func ClickHandler(w http.ResponseWriter, req *http.Request) {
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pbreq := &pb.Request{}
	err = protojson.Unmarshal(bytes, pbreq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = strategy.Reward(context.Background(), []bandits.ActionID{bandits.ActionID(pbreq.Uuid)}, &banditData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("data: %+v", banditData)
}
