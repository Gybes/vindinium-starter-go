package govindinium 

type State struct {
  Game Game_info
  Hero Hero_info
  Token string
  ViewUrl string
  PlayUrl string
}

type Game_info struct {
  Id string
  Turn int
  MaxTurns int
  Heroes []Hero_info
  Board Board_info
  Finished bool
}

type Board_info struct {
  Size int
  Tiles string
}

type Hero_info struct {
  Id int
  Name string
  UserId string
  Elo int
  Pos Pos_info
  Life int
  Gold int
  MineCount int
  SpawnPos Pos_info
  Crashed bool
}

type Pos_info struct {
  X int
  Y int
}
