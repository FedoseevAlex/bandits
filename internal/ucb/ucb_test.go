package ucb

import (
	"testing"
)

func TestUCB1(t *testing.T) {
	tests := []struct {
		name             string
		rewards          map[string]float64
		chosen           map[string]int
		totalRounds      int
		explorationRatio float64
		expected         string
		description      string
	}{
		{
			name:             "Single candidate",
			rewards:          map[string]float64{"A": 10.0},
			chosen:           map[string]int{"A": 20},
			totalRounds:      20,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should return the only candidate when there's only one option",
		},
		{
			name:             "Two candidates with different rewards",
			rewards:          map[string]float64{"A": 10.0, "B": 5.0},
			chosen:           map[string]int{"A": 10, "B": 10},
			totalRounds:      20,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should prefer candidate with higher reward",
		},
		{
			name:             "Multiple candidates with varying rewards",
			rewards:          map[string]float64{"A": 15.0, "B": 8.0, "C": 12.0, "D": 3.0},
			chosen:           map[string]int{"A": 18, "B": 30, "C": 22, "D": 30},
			totalRounds:      100,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should select candidate with highest UCB1 score",
		},
		{
			name:             "Zero rounds",
			rewards:          map[string]float64{"A": 5.0, "B": 3.0},
			chosen:           map[string]int{"A": 0, "B": 0},
			totalRounds:      0,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should handle zero total rounds",
		},
		{
			name:             "High rounds count",
			rewards:          map[string]float64{"A": 100.0, "B": 95.0},
			chosen:           map[string]int{"A": 500, "B": 500},
			totalRounds:      1000,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should work with high round counts",
		},
		{
			name:             "Zero rewards",
			rewards:          map[string]float64{"A": 0.0, "B": 0.0, "C": 0.0},
			chosen:           map[string]int{"A": 1, "B": 2, "C": 2},
			totalRounds:      5,
			explorationRatio: 1,
			expected:         "A", // Should return first one due to map iteration order
			description:      "Should handle zero rewards",
		},
		{
			name:             "Negative rewards",
			rewards:          map[string]float64{"A": -5.0, "B": -10.0, "C": -3.0},
			chosen:           map[string]int{"A": 40, "B": 40, "C": 20},
			totalRounds:      100,
			explorationRatio: 1,
			expected:         "C",
			description:      "Should handle negative rewards correctly",
		},
		{
			name:             "Very small rewards",
			rewards:          map[string]float64{"A": 0.001, "B": 0.002, "C": 0.0015},
			chosen:           map[string]int{"A": 30, "B": 30, "C": 30},
			totalRounds:      90,
			explorationRatio: 1,
			expected:         "B",
			description:      "Should work with very small reward values",
		},
		{
			name:             "Large reward differences",
			rewards:          map[string]float64{"A": 1000.0, "B": 1.0, "C": 500.0},
			chosen:           map[string]int{"A": 20, "B": 2, "C": 2},
			totalRounds:      50,
			explorationRatio: 1,
			expected:         "A",
			description:      "Should handle large differences in reward values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UCB1(tt.explorationRatio, tt.rewards, tt.chosen, tt.totalRounds)
			if result != tt.expected {
				t.Errorf("UCB1() = %v, want %v", result, tt.expected)
				t.Logf("Test description: %s", tt.description)
				t.Logf("Rewards: %v", tt.rewards)
				t.Logf("Total rounds: %d", tt.totalRounds)
			}
		})
	}
}

func TestUCB1ExplorationExploitation(t *testing.T) {
	// Test that UCB1 balances exploration and exploitation
	rewards := map[string]float64{
		"exploited": 10.0, // High reward, should be exploited
		"explored":  1.0,  // Low reward, but should be explored
	}

	// At low rounds, exploration term should dominate
	result := UCB1(1, rewards, map[string]int{"exploited": 19, "explored": 1}, 20)
	if result != "explored" {
		t.Errorf("Expected exploration at low rounds, got %s", result)
	}

	// At high rounds, exploitation term should dominate
	result = UCB1(1, rewards, map[string]int{"exploited": 500, "explored": 500}, 1000)
	if result != "exploited" {
		t.Errorf("Expected exploitation at high rounds, got %s", result)
	}
}

func TestUCB1MathematicalProperties(t *testing.T) {
	rewards := map[string]float64{"A": 10.0, "B": 5.0, "C": 8.0}
	chosen := map[string]int{"A": 10, "B": 10, "C": 10}
	totalRounds := 30

	// Test that the function is deterministic
	result1 := UCB1(1, rewards, chosen, totalRounds)
	result2 := UCB1(1, rewards, chosen, totalRounds)
	if result1 != result2 {
		t.Errorf("UCB1 should be deterministic, got %s and %s", result1, result2)
	}
}

func TestUCB1EdgeCases(t *testing.T) {
	// Test with very large numbers
	t.Run("Large numbers", func(t *testing.T) {
		candidates := map[string]float64{
			"A": 1e10,
			"B": 1e9,
		}
		result := UCB1(1, candidates, map[string]int{}, 1000)
		if result != "A" {
			t.Errorf("Expected A for highest reward, got %s", result)
		}
	})

	// Test with different data types
	t.Run("Integer candidates", func(t *testing.T) {
		candidates := map[int]float64{1: 10.0, 2: 5.0, 3: 8.0}
		result := UCB1(1, candidates, map[int]int{}, 10)
		if result != 1 {
			t.Errorf("Expected 1 for highest reward, got %d", result)
		}
	})
}

func TestUCB1Convergence(t *testing.T) {
	// Test that UCB1 converges to the best arm over time
	candidates := map[string]float64{
		"best":   10.0,
		"medium": 5.0,
		"worst":  1.0,
	}

	// Track selections over multiple rounds
	selections := make(map[string]int)
	for round := 1; round <= 100000; round++ {
		result := UCB1(1, candidates, selections, round)
		selections[result]++
	}

	// The best arm should be selected most frequently
	if selections["best"] < selections["medium"] || selections["medium"] < selections["worst"] {
		t.Errorf("Best arm should be selected most frequently. Selections: %v", selections)
	}
	t.Logf("Selections: %v", selections)
}

func BenchmarkUCB1(b *testing.B) {
	candidates := map[string]float64{
		"A": 10.0, "B": 5.0, "C": 8.0, "D": 3.0, "E": 7.0,
		"F": 12.0, "G": 4.0, "H": 9.0, "I": 6.0, "J": 2.0,
	}
	totalRounds := 100000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UCB1(1, candidates, map[string]int{}, totalRounds)
	}
}
