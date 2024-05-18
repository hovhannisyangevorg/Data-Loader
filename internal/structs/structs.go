package structs

type GlobalInfrastructura struct {
	NinjesAPIAnimal Animal
}

// Animal struct for each animal
type Animal struct {
	Name            string          `json:"name"`
	Taxonomy        Taxonomy        `json:"taxonomy"`
	Locations       []string        `json:"locations"`
	Characteristics Characteristics `json:"characteristics"`
}

// Taxonomy struct for taxonomy information
type Taxonomy struct {
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phylum"`
	Class          string `json:"class"`
	Order          string `json:"order"`
	Family         string `json:"family"`
	Genus          string `json:"genus"`
	ScientificName string `json:"scientific_name"`
}

// Characteristics struct for characteristics information
type Characteristics struct {
	Prey                string `json:"prey"`
	NameOfYoung         string `json:"name_of_young"`
	GroupBehavior       string `json:"group_behavior"`
	EstimatedPopSize    string `json:"estimated_population_size"`
	BiggestThreat       string `json:"biggest_threat"`
	DistinctiveFeature  string `json:"distinctive_feature"`
	GestationPeriod     string `json:"gestation_period"`
	LitterSize          string `json:"litter_size"`
	Habitat             string `json:"habitat"`
	Predators           string `json:"predators"`
	Diet                string `json:"diet"`
	Type                string `json:"type"`
	CommonName          string `json:"common_name"`
	NumberOfSpecies     string `json:"number_of_species"`
	Location            string `json:"location"`
	Slogan              string `json:"slogan"`
	Group               string `json:"group"`
	Color               string `json:"color"`
	SkinType            string `json:"skin_type"`
	TopSpeed            string `json:"top_speed"`
	Lifespan            string `json:"lifespan"`
	Weight              string `json:"weight"`
	Height              string `json:"height"`
	Length              string `json:"length"`
	AgeOfSexualMaturity string `json:"age_of_sexual_maturity"`
	AgeOfWeaning        string `json:"age_of_weaning"`
}
