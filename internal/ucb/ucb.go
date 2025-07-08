package ucb

import (
	"context"
	"math"

	"github.com/FedoseevAlex/bandits/internal/bandits"
)

func UCB1[T comparable](explorationRatio float64, rewards map[T]float64, chosen map[T]int, totalRounds int) T {
	var maxConfidence float64
	var candidateToShow T
	for candidate, candidateReward := range rewards {
		exploitation := candidateReward / float64(totalRounds+1)
		exploration := explorationRatio * math.Sqrt(2*math.Log(float64(totalRounds))/float64(chosen[candidate]+1))
		confidence := exploitation
		if !math.IsNaN(exploration) {
			confidence += exploration
		}
		if confidence > maxConfidence {
			candidateToShow = candidate
			maxConfidence = confidence
		}
	}
	return candidateToShow
}

type UCB1Strategy struct {
	explorationRatio float64
}

func (s *UCB1Strategy) Choose(ctx context.Context, data *bandits.ContextualData) (bandits.ActionID, error) {
	actionID := UCB1(s.explorationRatio, data.Rewards, data.Chosen, data.Rounds)
	data.Rounds++
	data.Chosen[actionID]++
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

func NewUCB1Strategy(explorationRatio float64) *UCB1Strategy {
	return &UCB1Strategy{
		explorationRatio: 0.1,
	}
}
