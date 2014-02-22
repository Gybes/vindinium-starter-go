package main

import (
  "github.com/aranasaurus/vindinium-starter-go/govindinium"
  "strconv"
  "flag"
  "fmt"
  "io/ioutil"
  "strings"
  "math/rand"
  "time"
)

func main() {
  flag.Parse()
  if (len(flag.Args()) < 2) {
    fmt.Println("Usage <[training|arena]> <number-of-games|number-of-turns> [server-url]")
  } else {
    args := flag.Args()
    mode := args[0]
    keyBuf, err := ioutil.ReadFile("secretKey")
    if (err != nil) {
      panic(err)
    }
    key := strings.Split(string(keyBuf), "\n")[0]

    var number_of_games, number_of_turns int
    var server_url string

    if (mode == "training") {
      number_of_games = 1
      number_of_turns, err = strconv.Atoi(args[1])
      if err != nil { panic(err.Error()) }
    } else {
      number_of_games, err = strconv.Atoi(args[1])
      if err != nil { panic(err.Error()) }
      number_of_turns = 300 //ignored in arena mode
    }

    if (len(args) == 3) {
      server_url = args[2]
    } else {
      server_url = "http://vindinium.org"
    }

    for i:= 0; i<number_of_games; i++ {
      govindinium.Start(server_url, key, mode, number_of_turns, NewRandomBot())
      fmt.Printf("Game finished: %d/%d\n", i+1, number_of_games)
    }
  }
}

/*****************
*** Random Bot ***
*****************/

type RandomBot struct {
  Dirs []string
}

func NewRandomBot() *RandomBot {
  return &RandomBot{Dirs: []string{ "Stay", "North", "South", "East", "West" }}
}

func (r *RandomBot)Move(state govindinium.State) (string) {
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

func (s *SlowBot)Move(state govindinium.State) (string) {
  time.Sleep(2000)
  return s.Dirs[rand.Intn(len(s.Dirs))]
}
