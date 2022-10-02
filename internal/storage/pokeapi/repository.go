package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	pokemoncli "pokemon_http_response/internal"
)

const (
    getEndpoint = "/pokemon"
    apiUrl      = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
    url string
}

func NewPokeapiRepository() pokemoncli.PokemonRepo {
    return &pokemonRepo{url: apiUrl}
}

func (p *pokemonRepo) GetPokemons(count int) (pokemons []pokemoncli.Pokemon, err error) {
    url := fmt.Sprintf("%v%v?limit=%v&offset=0", p.url, getEndpoint, count)
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var res pokemoncli.Response
    err = json.Unmarshal(contents, &res)
    if err != nil {
        return nil, err
    }
    pokemons = res.Pokemons
    return
}