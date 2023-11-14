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

### Problem CRUD

#### Get active problems  

*given*: user wants to fetch active problems  
*when*: issues a request to api  
*then*: problem instances are returned from the storage  

#### Get all problems even deleted ones  

*given*: user wants to fetch all problems  
*when*: issues a request to api with query parameter deleted=true  
*then*: all problem instances are returned from the storage  

#### Create problem  

*given*: user wants to create a problem  
*when*: issues a request to api with problem_id and problem_description  
*then*: problem instance is created and added to storage without any actions added  

#### Delete problem  

*given*: user wants to delete a problem  
*when*: issues a request to api with problem_id  
*then*: problem instance marked as deleted and won't be shown  
