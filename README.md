# Bandits

This project is intented to be the implementation of multi-armed (MAB) bandit framework.  
I want to try make an app that would be easy to use if you need MAB.

## Glossary

| Term    | Meaning                                                                                 | Example                                                   |
| ------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| MAB     | Very powerful framework for algorithms that make decisions over time under uncertainty. |                                                           |
| Problem | Problem which we want to solve by using MAB framework.                                  | Increase revenue from ads.                                |
| Action  | One of the possible decisions which MAB can choose.                                     | Ad or news story to show.                                 |
| Reward  | This is a reward that is given for actions if they succeed.                             | If ad is clicked then its reward is +1 click.             |
| Context | Some environment MAB should consider making a decision.                                 | Age of the user to consider when choosing an ad for them. |

## Planned features and scenarios

There will be two APIs(?):

1. Strict API with entity hierarchy

   - One will need to create explicit problem, attach actions and start experiment.
   - IDs of problem and actions should be fixed until the end of experiment.
   - Actions can be added on the fly, but their performance will be up to MAB strategy (UCB, E-greedy, etc).
   - If context is specified then it will affect the stats.

   GET POST /problems
   GET POST /actions
   POST /upload_data
   GET POST /choose
   GET POST /reward

2. Loose API

   - All you need is to give a list of action IDs to choose from.
   - Actions will be automatically created and accounted.
   - No context support.

   GET POST /choose
   GET POST /reward
   POST /upload_data

I need to think about how it would be integrated into the real set ups.  
Based on that API will be formalized.

### Typical scenarios

#### Get active problems

_given_: user wants to fetch active problems  
_when_: issues a request to api  
_then_: problem instances are returned from the storage

GET /problems

#### Get all problems even deleted ones

_given_: user wants to fetch all problems  
_when_: issues a request to api with query parameter deleted=true  
_then_: all problem instances are returned from the storage

GET /problems?all=true

#### Create problem

_given_: user wants to create a problem  
_when_: issues a request to api with `problem_id` and `problem_description`
_then_: problem instance is created and added to storage without any actions added

POST /problems
{
"name": "problem_name",
""
}

#### Delete problem

_given_: user wants to delete a problem  
_when_: issues a request to api with `problem_id`  
_then_: problem instance marked as deleted and won't be shown
