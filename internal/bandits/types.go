package bandits

import "context"

type MABContext string

// Should usually start with action_
type ActionID string

// Should usually start with problem_
type ProblemID string

/*
This is an interface that represents any possible MAB strategy (UCB, Epsilon-Greedy, etc..)
Basically every strategy needs a MAB context and data
*/
type Strateger interface {
	Choose(ctx context.Context, data *ContextualData) (ActionID, error)
	Reward(ctx context.Context, actions []ActionID, data *ContextualData) error
}

type Storager interface {
	GetData(ctx context.Context) (Data, error)
	SetData(ctx context.Context, data Data) error
}

// Problem is a struct that represents a MAB problem.
// For example, in a ad campaign, a problem could be to choose the best ad to show to a user.
// This struct represents the entity that is being optimized.
type Problem struct {
	ID          ProblemID
	Description string
	Actions     []*Action
	Storage     Storager
	Strategy    Strateger
}

// Action is a struct that represents an action that can be chosen by the MAB strategy.
// For example, in a recommendation system, an action could be a product, a movie, a news article, etc.
// Or in a ad system, an action could be an ad, a banner, a sponsored post, etc.
type Action struct {
	ID          ActionID
	Description string
	Payload     []byte
}

type ContextualData struct {
	Rounds  int
	Rewards map[ActionID]float64
}

type Data map[MABContext]ContextualData
