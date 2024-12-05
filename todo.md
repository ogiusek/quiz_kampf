done:
1. repair hw query parsing
2. add change question time command
struct {
  q_id
  seconds
}
3. add parsing questions to dtos
todo:
4. start server where:
- send match state (without answers)
- on start you wait 3 seconds                     // use cases 
- send question                                   // use cases
- wait for response                               // use cases 
- on response {
  send everyone who answered and do correct and his answer
}
- if not responded after answer time end question // use cases
- if questions left go to step 1                  // use cases
- else {
  send match state
}

so you have to implement:
### messages:
- in game match state
- answered 
- question

### handlers:
- start
- answer (use mutex so only one person at a time can answer)