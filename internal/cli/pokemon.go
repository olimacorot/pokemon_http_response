package cli

import (
	"fmt"
	pokemoncli "pokemon_http_response/internal"
	"pokemon_http_response/internal/output"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, arg []string)

const (
    countFlag  = "count"
    exportFlag = "csv"
    nameFlag   = "name"
)

func InitPokemonCmd(repository pokemoncli.PokemonRepo) *cobra.Command {
    pokemonCmd := &cobra.Command{
        Use:   "pokemons",
        Short: "Get list of the pokemon",
        Run:   runPokemonFn(repository),
    }

    pokemonCmd.Flags().IntP(countFlag, "c", 10, "Total items of the list")
    pokemonCmd.Flags().Bool(exportFlag, false, "Export to csv - default: false.")
    pokemonCmd.Flags().StringP(nameFlag, "n", "Pokemons", "Asing name to the new file")

    return pokemonCmd
}

func runPokemonFn(repository pokemoncli.PokemonRepo) CobraFn {
    return func(cmd *cobra.Command, args []string) {
        count, _ := cmd.Flags().GetInt(countFlag)
        export, _ := cmd.Flags().GetBool(exportFlag)
        fileName, _ := cmd.Flags().GetString(nameFlag)

        pokemons, _ := repository.GetPokemons(count)

        if export {
            g := output.Outputter{pokemons, fileName}
            err := g.ExportCsv()
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Printf("Your new %v.csv was exported successfully.", fileName)
            }
        } else {
            fmt.Println(pokemons)
        }
    }
}


