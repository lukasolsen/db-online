package schemas

type SuperType string

const (
	Pokemon SuperType = "Pokémon"
	Trainer SuperType = "Trainer"
	Energy  SuperType = "Energy"
)

type PokemonAbility struct {
	name  string `json:"name"`
	text  string `json:"text"`
	_type string `json:"type"`
}

type PokemonAttack struct {
	name                string   `json:"name"`
	cost                []string `json:"cost"`
	convertedEnergyCost int      `json:"convertedEnergyCost"`
	damage              string   `json:"damage"`
	text                string   `json:"text"`
}

type PokemonWeaknessSchema struct {
	_type string `json:"type"`
	value string `json:"value"`
}

type PokemonLegalitiesSchema struct {
	unlimited string `json:"unlimited"`
}

type PokemonSchema struct {
	id                     string                  `json:"id"`
	name                   string                  `json:"name"`
	superType              SuperType               `json:"supertype"`
	subTypes               []string                `json:"subtypes"`
	level                  string                  `json:"level"`
	hp                     string                  `json:"hp"`
	types                  []string                `json:"types"`
	evolvesFrom            string                  `json:"evolves_from"`
	abilities              []PokemonAbility        `json:"abilities"`
	attacks                []PokemonAttack         `json:"attacks"`
	weaknesses             []PokemonWeaknessSchema `json:"weaknesses"`
	retreatCost            []string                `json:"retreat_cost"`
	convertedRetreatCost   string                  `json:"converted_retreat_cost"`
	number                 string                  `json:"number"`
	artist                 string                  `json:"artist"`
	rarity                 string                  `json:"rarity"`
	flavorText             string                  `json:"flavor_text"`
	nationalPokedexNumbers []string                `json:"national_pokedex_numbers"`
	legalities             PokemonLegalitiesSchema `json:"legalities"`
	images                 []struct {
		small string `json:"small"`
		large string `json:"large"`
	} `json:"images"`
}

type PokemonSchemaResponse struct {
	cardId                 string `json:"card_id"`
	name                   string `json:"name"`
	superType              string `json:"supertype"`
	subTypes               string `json:"subtypes"`
	level                  string `json:"level"`
	hp                     string `json:"hp"`
	types                  string `json:"types"`
	evolvesFrom            string `json:"evolves_from"`
	retreatCost            string `json:"retreat_cost"`
	convertedRetreatCost   string `json:"converted_retreat_cost"`
	number                 string `json:"number"`
	artist                 string `json:"artist"`
	rarity                 string `json:"rarity"`
	flavorText             string `json:"flavor_text"`
	nationalPokedexNumbers string `json:"national_pokedex_numbers"`
	images                 string `json:"images"`
	abilities              string `json:"abilities"`
	attacks                string `json:"attacks"`
	weaknesses             string `json:"weaknesses"`
	legalities             string `json:"legalities"`
}
