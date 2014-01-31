package govindinium

import (
  "strconv"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "fmt"
)

type StatusCodeError struct {
  s string
}

func (e *StatusCodeError)Error() string {
  return e.s
}

func Json_encode(resp *http.Response) (State, error) {
  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil { return State{}, err }

    var data State
    err = json.Unmarshal(body, &data)
    if err != nil {
      fmt.Printf("%T\n%s\n%#v\n",err, err, err)
      panic(err.Error())
    }

    return data, nil
  } else {
    return State{}, &StatusCodeError{strconv.Itoa(resp.StatusCode)}
  }
}

func Get_new_game_state(server_url string, key string, mode string, turns int) (State, error) {
  v := url.Values{}
  var api_endpoint string

  if mode == "training" {
    v.Set("key", key)
    v.Set("turns", strconv.Itoa(turns))
    v.Set("map", "m1")

    api_endpoint = "/api/training"
  }else {
    v.Set("key", key)

    api_endpoint = "/api/arena"
  }

  resp, err := http.PostForm(server_url + api_endpoint, v)
  if err != nil { return State{}, err }

  return Json_encode(resp)
}

func Move(play_url string, direction string) (State, error) {
  resp, err := http.PostForm(play_url, url.Values{ "dir" : { direction }})
  if err != nil { return State{}, err }

  return Json_encode(resp)
}

func Is_finished(s State) (bool) {
  return s.Game.Finished
}

func Start(server_url string, key string, mode string, turns int, bot Bot) {
  if mode == "arena" {
    fmt.Println("Connected and waiting for other players to join...")
  }
  state,err := Get_new_game_state(server_url, key, mode, turns)
  if err != nil { panic(err.Error()) }

  fmt.Println("Playing at: ", state.ViewUrl)

  for !Is_finished(state) {
    fmt.Print(".")

    url := state.PlayUrl
    direction := bot.Move(state)
    state, err = Move(url, direction)
    if err != nil { panic(err.Error()) }
  }

  fmt.Println()
}
