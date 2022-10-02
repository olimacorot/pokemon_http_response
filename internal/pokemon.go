package pokemoncli

type Response struct {
    Pages    int       `json:"count"`
    Next     string    `json:"next"`
    Previous string    `json:"previous"`
    Pokemons []Pokemon `json:"results"`
}

type Pokemon struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type PokemonRepo interface {
    GetPokemons(count int) ([]Pokemon, error)
}
