package structs

type AnimalOperation struct {
	AnimalName []string
}

func NewAnimalOperation() *AnimalOperation {
	return &AnimalOperation{
		AnimalName: []string{
			"dog", //"cat", "horse", //"cow", "sheep", "pig", "goat", "chicken", "duck", "turkey",
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

// 			"cheetah", "lion", "tiger", "elephant", "zebra", "giraffe", "penguin", "koala", "panda", "bear", "wolf", "rhinoceros", "hippopotamus", "crocodile", "snake", "frog", "ostrich", "gorilla", "monkey", "dolphin", "shark", "whale", "octopus", "seahorse", "turtle", "parrot", "eagle", "owl", "bat", "deer", "rabbit", "horse", "cow", "goat", "sheep", "pig", "chicken", "duck", "goose", "swan", "peacock", "flamingo", "pigeon", "cockatoo", "macaw", "hummingbird", "blue jay", "cardinal", "woodpecker", "kingfisher", "puffin", "albatross", "seagull", "pelican", "stork", "ibis", "heron", "crane", "egret", "sandpiper", "curlew", "avocet", "phalarope", "tern", "gull", "skua", "plover", "lapwing", "sandpiper", "godwit", "whimbrel", "snipe", "curlew", "woodcock", "rail", "coot", "moorhen", "crake", "gallinule", "swamphen", "bustard", "crane", "bittern", "flamingo", "ibis", "spoonbill", "stork", "pelican", "albatross", "petrel", "storm petrel", "shearwater", "gannet", "frigatebird", "penguin", "booby", "skua", "jaeger", "gull", "tern", "skimmer", "auklet", "murre", "guillemot", "puffin", "grebe", "albatross", "fulmar", "tern", "petrel", "shearwater", "storm petrel", "gannet", "pelican", "booby", "cormorant", "frigatebird", "penguin", "skua", "jaeger", "gull", "tern", "skimmer", "auklet", "murre", "guillemot", "puffin", "grebe", "albatross", "fulmar", "stork", "ibis", "spoonbill", "heron", "egret", "bittern", "flamingo", "swan", "goose", "duck", "teal", "mallard", "pintail", "gadwall", "wigeon", "shoveler", "scaup",

//func NewAnimalOperation() *AnimalOperation {
//	return &AnimalOperation{
//		AnimalName: []string{
//			//"fox", "cheetah", "lion", "tiger", "elephant", "zebra", "giraffe", "penguin", "koala", "panda", "bear", "wolf", "rhinoceros", "hippopotamus", "crocodile", "snake", "frog", "ostrich", "gorilla", "monkey", "dolphin", "shark", "whale", "octopus", "seahorse", "turtle", "parrot", "eagle", "owl", "bat", "deer", "rabbit", "horse", "cow", "goat", "sheep", "pig", "chicken", "duck", "goose", "swan", "peacock", "flamingo", "pigeon", "cockatoo", "macaw", "hummingbird", "sparrow",
//			//"cardinal", "woodpecker", "kingfisher", "puffin", "albatross", "seagull",
//			//"pelican", "stork", "ibis", "heron", "crane", "egret",
//			//"avocet", "tern", "gull",
//			//"coot", "moorhen", "bustard", "crane", "flamingo", "ibis", "spoonbill", "stork", "pelican", "albatross",
//			"frigatebird", "penguin", "booby", "skua", "jaeger", "gull", "tern", "skimmer", "auklet", "murre", "guillemot", "puffin", "grebe", "albatross", "fulmar", "tern", "petrel", "shearwater", "storm petrel", "gannet", "pelican", "booby", "cormorant", "frigatebird", "penguin", "skua", "jaeger", "gull", "tern", "skimmer", "auklet", "murre", "guillemot", "puffin", "grebe", "albatross", "fulmar", "stork", "ibis", "spoonbill", "heron", "egret", "bittern", "flamingo", "swan", "goose", "duck", "teal", "mallard", "pintail", "gadwall", "wigeon", "shoveler", "scaup",
//		},
//	}
//}
