module github.com/yash-278/pokedexcli

go 1.20

require github.com/yash-278/pokedexcli/internal/pokedexapi v1.0.0
replace github.com/yash-278/pokedexcli/internal/pokedexapi => ./internal/pokedexapi


require github.com/yash-278/pokedexcli/internal/pokedex_cache v1.0.0
replace github.com/yash-278/pokedexcli/internal/pokedex_cache => ./internal/pokedex_cache