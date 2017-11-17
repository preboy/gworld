package session

import (
	"fmt"
)

type StageFunction struct {
	OnEnter  func(s *Session)
	OnLeave  func(s *Session)
	OnUpdate func(s *Session)
}

var (
	stages [10]StageFunction
	idx    int
)

func init() {

	idx = 0 // login
	stages[idx].OnEnter = func(s *Session) {
	}
	stages[idx].OnLeave = func(s *Session) {
	}
	stages[idx].OnUpdate = func(s *Session) {
	}

	idx = 1 // login
	stages[idx].OnEnter = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnEnter")
	}

	stages[idx].OnLeave = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnLeave")
	}

	stages[idx].OnUpdate = func(s *Session) {
		// fmt.Println("Stage:", s.stage_id, "OnUpdate")
	}

	idx = 2 // game
	stages[idx].OnEnter = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnEnter")
	}

	stages[idx].OnLeave = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnLeave")
	}

	stages[idx].OnUpdate = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnUpdate")
	}
}

func Next(s *Session) {
	stages[s.stage_id].OnLeave(s)
	s.stage_id++
	stages[s.stage_id].OnEnter(s)
}
