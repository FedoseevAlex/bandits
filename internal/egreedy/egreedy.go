package egreedy

import (
	"context"
	"math/rand/v2"

	"github.com/FedoseevAlex/bandits/internal/bandits"
)

type _EGreedyStrategy struct {
	epsilon float64
}

func _EGreedy[T comparable](epsilon float64, rewards map[T]float64) T {
	var bestAction T
	actions := make([]T, 0, len(rewards))
	maxReward := float64(-1)
	for key := range rewards {
		if rewards[key] > maxReward {
			maxReward = rewards[key]
			bestAction = key
		}
		actions = append(actions, key)
	}
	if rand.Float64() < epsilon {
		// Explore: choose a random action
		return actions[rand.IntN(len(actions))]
	}
	return bestAction
}

func NewEGreedyStrategy(epsilon float64) *_EGreedyStrategy {
	return &_EGreedyStrategy{epsilon: epsilon}
}

func (s *_EGreedyStrategy) Choose(ctx context.Context, data *bandits.ContextualData) (bandits.ActionID, error) {
	return _EGreedy(s.epsilon, data.Rewards), nil
}

func (s *_EGreedyStrategy) Reward(ctx context.Context, actions []bandits.ActionID, data *bandits.ContextualData) error {
	// Simple reward function that gives 1 reward to each specified action
	for _, actionID := range actions {
		data.Rewards[actionID] = data.Rewards[actionID] + 1
	}
	return nil
}
