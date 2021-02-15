package interfaces

type Handler interface {
	Do(*Contexter)
}
