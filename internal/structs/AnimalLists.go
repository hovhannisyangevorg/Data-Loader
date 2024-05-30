package structs

type AnimalOperation struct {
	AnimalName []string
}

func NewAnimalOperation() *AnimalOperation {
	return &AnimalOperation{
		AnimalName: []string{
			"dog", "cat", "horse", // "cow", "sheep", "pig", "goat", "chicken", "duck", "turkey",
			//"elephant", "lion", "tiger", "leopard", "cheetah", "giraffe", "zebra", "hippopotamus",
			//"rhinoceros", "buffalo", "kangaroo", "koala", "panda", "gorilla", "chimpanzee",
			//"wolf", "fox", "bear", "deer", "moose", "elk", "buffalo", "bison", "hare", "rabbit",
			//"rat", "mouse", "squirrel", "chipmunk", "beaver", "otter", "badger", "raccoon", "skunk",
			//"weasel", "ferret", "mink", "lynx", "bobcat", "cougar", "panther", "jaguar", "ocelot",
			//"caracal", "lynx", "puma", "cougar", "leopard", "ocelot", "jaguarundi", "tiger", "lion", "cheetah",
			//"serval", "caracal", "ocelot", "wolf", "dingo", "coyote", "jackal", "dhole",
			//"fox", "peccary", "hippopotamus", "camel", "llama", "alpaca", "giraffe",
			//"okapi", "deer", "elk", "moose", "caribou", "reindeer", "muntjac", "whale", "dolphin", "porpoise", "narwhal",
			//"seal", "walrus", "sloth", "anteater", "armadillo", "aardvark", "elephant", "rhinoceros", "hippopotamus",
			//"tapir", "manatee", "dugong", "hyrax", "chinchilla", "capybara", "paca",
			//"hamster", "gerbil", "rat", "mouse", "lemming", "vole", "muskrat", "beaver", "squirrel",
			//"marmot", "chipmunk", "gopher", "groundhog", "porcupine", "hedgehog", "shrew", "mole",
			//"stingray", "skate", "sawfish", "carp", "catfish", "salmon", "trout", "pike", "walleye", "carp", "goldfish",
			//"koi", "tuna", "herring", "sardine", "marlin", "shark",
			//"ray", "stingray", "skate", "grouper", "bass", "cod", "halibut", "flounder",
			//"sunfish", "crappie", "pufferfish", "boxfish", "flounder", "halibut", "haddock", "cod", "ling", "herring", "sardine", "tuna", "bonito",
			//"marlin", "barracuda", "wahoo", "kingfish", "grouper", "bass", "cod", "haddock", "ling", "haddock",
			//"anglerfish", "monkfish", "stargazer", "toadfish", "catfish", "loach", "eel", "lamprey", "hagfish", "shark",
			//"stingray", "skate",
		},
	}
}
