package egreedy

import (
	"testing"

	"github.com/FedoseevAlex/bandits/internal/bandits"
)

func TestEGreedy(t *testing.T) {
	numberOfRounds := 1000
	epsilon := 0.1
	rewards := map[bandits.ActionID]float64{
		"A": 10.0,
		"B": 5.0,
		"C": 8.0,
	}

	chosen := make(map[bandits.ActionID]int)
	for range numberOfRounds {
		result := _EGreedy(epsilon, rewards)
		chosen[result]++
	}

	if chosen["A"] < chosen["B"] || chosen["A"] < chosen["C"] {
		t.Errorf("Best action should be chosen most frequently. Selections: %v", chosen)
	}
}
