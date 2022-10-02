package output

import (
	"encoding/csv"
	"fmt"
	"os"
	pokemoncli "pokemon_http_response/internal"
)

type Outputter struct {
    Pokemons []pokemoncli.Pokemon
    FileName  string
}

func (out Outputter) ExportCsv() error {
    file, err := os.Create(fmt.Sprintf("%v.csv", out.FileName))

    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range out.Pokemons {
        data := [] string {value.Name, value.Url}
        if err := writer.Write(data); err != nil {
            return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
        }
    }
    return nil
}