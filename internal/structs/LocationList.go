package structs

type LocationOperation struct {
	LocationName []string
}

func NewLocationOperation() *LocationOperation {
	return &LocationOperation{
		LocationName: []string{"Africa", "Antarctica", "Asia", "Central-America", "Eurasia", "Europe", "North-America", "Oceania", "South-America"},
	}
}
