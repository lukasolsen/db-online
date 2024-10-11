package updater

import (
	"fmt"

	"github.com/codevault-llc/db-online/pkg/lib"
	"github.com/codevault-llc/db-online/pkg/schemas"
)

func updatePokemon() {
	client := lib.NewClient()

	// Fetch the data from the API
	data, err := client.FetchURL("https://api.pokemontcg.io/v1/cards", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(data))

	// Unmarshal the data into the PokemonSchemaResponse struct
	var response schemas.PokemonSchemaResponse
	err = client.ToJSON(data, &response)

	if err != nil {
		panic(err)
	}

}
