package Models

type Status int

const (
	Empty = 0
	InUse
	ReadyForPickUp
)

type Cubby struct {
	Id           int
	Status       Status
	CurrentOrder Order
}
