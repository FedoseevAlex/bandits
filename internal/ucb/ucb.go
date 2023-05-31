package ucb

import (
	"context"
	"math"

	"github.com/FedoseevAlex/bandits/internal/bandits"
)

func UCB1[T comparable](candidates map[T]float64, totalRounds int) T {
	var maxConfidence float64
	var candidateToShow T
	for candidate, candidateReward := range candidates {
		exploitation := candidateReward / float64(totalRounds+1)
		exploration := math.Sqrt(2 * math.Log(float64(totalRounds)) / (candidateReward + 1))
		confidence := exploration + exploitation
		if confidence >= maxConfidence {
			candidateToShow = candidate
			maxConfidence = confidence
		}
	}
	return candidateToShow
}

type UCB1Strategy struct {
}

func (s *UCB1Strategy) Choose(ctx context.Context, data *bandits.ContextualData) (bandits.ActionID, error) {
	actionID := UCB1(data.Rewards, data.Rounds)
	data.Rounds++
	return actionID, nil
}

func (s *UCB1Strategy) Reward(ctx context.Context, actions []bandits.ActionID, data *bandits.ContextualData) error {
	// Simple reward function that gives 1 reward to each specified action
	for _, actionID := range actions {
		data.Rewards[actionID] = data.Rewards[actionID] + 1
	}
	return nil
}

var _ bandits.Strateger = &UCB1Strategy{}

func NewUCB1Strategy() *UCB1Strategy {
	return &UCB1Strategy{}
}
