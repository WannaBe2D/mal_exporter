package models

type Meta struct {
	Name string `json:"name"`
}

type AnimeMini struct {
	Episode int  `json:"increment"`
	Meta    Meta `json:"meta"`
}

type AnimeExel struct {
	ID         int    `json:"mal_id"`
	Name       string `json:"title"`
	Type       string `json:"type"`
	Score      int    `json:"score"`
	Wepiosedes int    `json:"watched_episodes"`
	Tepisodes  int    `json:"total_episodes"`
	Url        string `json:"url"`
}

type WatchlistExel struct {
	Anime []AnimeExel `json:"anime"`
}

type History struct {
	History []AnimeMini `json:"history"`
}

type Studio struct {
	ID   int    `json:"mal_id"`
	Name string `json:"name"`
}

type Genre struct {
	ID   int    `json:"mal_id"`
	Name string `json:"name"`
}

type Anime struct {
	ID        int      `json:"mal_id"`
	Name      string   `json:"title"`
	Wepisodes int      `json:"watched_episodes"`
	Aepisodes int      `json:"total_episodes"`
	Image     string   `json:"image_url"`
	Genres    []Genre  `json:"genres"`
	Studios   []Studio `json:"anime_studios"`
}

type WatchListJson struct {
	Anime []Anime `json:"anime"`
}
