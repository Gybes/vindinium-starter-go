package main

import (
  "github.com/Gybes/vindinium-starter-go"
  "strconv"
  "flag"
  "fmt"
)

func main() {
  flag.Parse()
  if (len(flag.Args()) < 3) {
    fmt.Println("Usage <key> <[training|arena]> <number-of-games|number-of-turns> [server-url]")
  }else {
    args := flag.Args()
    key := args[0]
    mode := args[1]

    var number_of_games, number_of_turns int
    var server_url string
    var err error

    if(mode == "training") {
      number_of_games = 1
      number_of_turns, err = strconv.Atoi(args[2])
      if err != nil { panic(err.Error()) }
    }else {
      number_of_games, err = strconv.Atoi(args[2])
      if err != nil { panic(err.Error()) }
      number_of_turns = 300 //ignored in arena mode
    }

    if(len(args) == 4) {
      server_url = args[3]
    }else {
      server_url = "http://vindinium.org"
    }

    for i:= 0; i<number_of_games; i++ {
      govindinium.Start(server_url, key, mode, number_of_turns, govindinium.NewRandomBot())
      fmt.Printf("Game finished: %d/%d\n", i+1, number_of_games)
    }
  }
}
