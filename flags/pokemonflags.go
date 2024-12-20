package flags

import (
	"flag"
	"os"

	"github.com/hazeliscoding/godex/connections"
)

func SetupPokemonFlagSet() (*flag.FlagSet, *bool) {
	pokeFlags := flag.NewFlagSet("pokeFlags", flag.ExitOnError)

	typesFlag := pokeFlags.Bool("types", false, "Print the declared Pokémon's typing")

	return pokeFlags, typesFlag
}

func TypesFlag() error {
	pokemonName := os.Args[1]

	connections.TypeApiCall(pokemonName, "https://pokeapi.co/api/v2/pokemon/")

	return nil
}
