# Learning about web server from nginx

## Questions
[Questions when I read nginx guide book](./question.md).

## Code analysis
### 1. State Machine for Http analyzing
- What is State Machine?  
First we google it and found a [page](https://www.itemis.com/en/yakindu/state-machine/documentation/user-guide/overview_what_are_state_machines) about it. 
Keywords:  
    - Finite number of state, also called FSM.
    - Based on the current state and a given input the machine performs state transitions and produces outputs.
There are two basic main FSM from automatic theory, Moore and Mealy.  
    - Moore
    States are able to produce outputs, and the output is determined solely by the current state, not by any input.  
    - Mealy
    Mealy machines produce outputs only on transitions and not in states.  
    