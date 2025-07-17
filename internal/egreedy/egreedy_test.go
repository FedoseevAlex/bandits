package egreedy

import (
	"math/rand/v2"
	"testing"

	"github.com/FedoseevAlex/bandits/internal/bandits"
)

func TestEGreedyStatic(t *testing.T) {
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

func TestEgreedyDynamic(t *testing.T) {
	numberOfRounds := 100000
	epsilon := 0.3
	rewards := map[bandits.ActionID]float64{
		"A": 0,
		"B": 0,
		"C": 0,
	}
	conversions := map[bandits.ActionID]float64{
		"A": 0.5,
		"B": 0.01,
		"C": 0.01,
	}

	chosen := make(map[bandits.ActionID]int)
	for range numberOfRounds {
		result := _EGreedy(epsilon, rewards)
		chosen[result]++
		if rand.Float64() < conversions[result] {
			rewards[result] += 1
		}
	}

	if chosen["A"] < chosen["B"] || chosen["A"] < chosen["C"] {
		t.Errorf("Best action should be chosen most frequently. Selections: %v, rewards: %v", chosen, rewards)
	}

}
