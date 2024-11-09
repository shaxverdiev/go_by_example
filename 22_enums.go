package main

import "fmt"

type ServerState int


const (
	StateIdle ServerState = iota  // 0
	StateConnected                // 1
	StateError                    // 2
	StateRetrying                 // 3
)

var stateName = map[ServerState]string{
	StateIdle:	     "idle",
	StateConnected:  "connected",
	StateError:      "error",
	StateRetrying:   "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]            // 2) и будет реализована эта логика
}

func main() {
	// ns := transition(StateIdle)
	// fmt.Println(ns)
	fmt.Println(StateConnected, "<") // 1) эта строка выведет не число 1, а "connected" тк тип ServerState реализует интерфейс fmt.Stringer (тип ServerState имеет метод String())

	// ns2 := transition(ns)
	// fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unlnow state: %s", s))
	}
}