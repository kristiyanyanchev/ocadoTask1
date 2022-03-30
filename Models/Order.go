package Models

type Order struct {
	Id            int
	Items         []Item
	AwaitingItems []Item
}
