package main

import (
	pokemoncli "pokemon_http_response/internal"
	"pokemon_http_response/internal/cli"
	"pokemon_http_response/internal/storage/pokeapi"

	"github.com/spf13/cobra"
)

func main() {

    var repo pokemoncli.PokemonRepo
    repo = pokeapi.NewPokeapiRepository()

    rootCmd := &cobra.Command{Use: "pokemons-cli"}
    rootCmd.AddCommand(cli.InitPokemonCmd(repo))
    rootCmd.Execute()
}