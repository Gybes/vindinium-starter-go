package govindinium

type Bot interface {
  Move(state State) string
}
