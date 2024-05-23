package structs

type LocationOperation struct {
	LocationName []string
}

func NewLocationOperation() *LocationOperation {
	return &LocationOperation{
		LocationName: []string{},
	}
}
