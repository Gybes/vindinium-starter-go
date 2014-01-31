package govindinium

import (
  "math/rand"
  "time"
)

/*****************
*** Random Bot ***
*****************/

type Bot interface {
  Move(state State) string
}

type RandomBot struct {
  Dirs []string
}

func NewRandomBot() *RandomBot {
  return &RandomBot{Dirs: []string{ "Stay", "North", "South", "East", "West" }}
}

func (r *RandomBot)Move(state State) (string) {
  return r.Dirs[rand.Intn(len(r.Dirs))]
}

/***************
*** Slow Bot ***
***************/

type SlowBot struct {
  Dirs []string
}

func NewSlowBot() *SlowBot {
  return &SlowBot{Dirs: []string{ "Stay", "North", "South", "East", "West" }}
}

func (s *SlowBot)Move(state State) (string) {
  time.Sleep(2000)
  return s.Dirs[rand.Intn(len(s.Dirs))]
}
