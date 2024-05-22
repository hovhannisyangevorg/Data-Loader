package structs

type GlobalInfrastructura struct {
	Animal []Animal
}

type (
	// Animal struct for each animal
	Animal struct {
		Name            string          `json:"name"`
		Taxonomy        Taxonomy        `json:"taxonomy"`
		Locations       []string        `json:"locations"`
		Characteristics Characteristics `json:"characteristics"`
	}
	// Taxonomy struct for taxonomy information
	Taxonomy struct {
		Kingdom        string `json:"kingdom"`
		Phylum         string `json:"phylum"`
		Class          string `json:"class"`
		Order          string `json:"order"`
		Family         string `json:"family"`
		Genus          string `json:"genus"`
		ScientificName string `json:"scientific_name"`
	}
	// Characteristics struct for characteristics information
	Characteristics struct {
		Diet                    string `json:"diet,omitempty"`
		AverageLitterSize       string `json:"average_litter_size,omitempty"`
		Type                    string `json:"type"`
		CommonName              string `json:"common_name,omitempty"`
		Slogan                  string `json:"slogan,omitempty"`
		Group                   string `json:"group,omitempty"`
		Color                   string `json:"color"`
		SkinType                string `json:"skin_type"`
		Lifespan                string `json:"lifespan"`
		Weight                  string `json:"weight,omitempty"`
		Length                  string `json:"length,omitempty"`
		Height                  string `json:"height,omitempty"`
		TopSpeed                string `json:"top_speed,omitempty"`
		NameOfYoung             string `json:"name_of_young,omitempty"`
		GroupBehavior           string `json:"group_behavior,omitempty"`
		EstimatedPopulationSize string `json:"estimated_population_size,omitempty"`
		BiggestThreat           string `json:"biggest_threat,omitempty"`
		MostDistinctiveFeature  string `json:"most_distinctive_feature,omitempty"`
		Habitat                 string `json:"habitat,omitempty"`
		Predators               string `json:"predators,omitempty"`
		Lifestyle               string `json:"lifestyle,omitempty"`
		FavoriteFood            string `json:"favorite_food,omitempty"`
		MainPrey                string `json:"main_prey,omitempty"`
		OtherNames              string `json:"other_name,omitempty"`
		GestationPeriod         string `json:"gestation_period,omitempty"`
		LitterSize              string `json:"litter_size,omitempty"`
		AgeOfSexualMaturity     string `json:"age_of_sexual_maturity,omitempty"`
		AgeOfWeaning            string `json:"age_of_weaning,omitempty"`
		NumberOfSpecies         string `json:"number_of_species,omitempty"`
	}
)
