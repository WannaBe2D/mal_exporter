package controllers

import (
	"exporter/services"
	"fmt"
)

type WatchlistController struct {
}

func (w *WatchlistController) Create(username string) {

	watchlist := new(services.WatchlistService)

	body, err := watchlist.ParseWatchlist(username)

	if err != nil {
		return
	}

	watchlist.CreateJson(body)
}

func (w *WatchlistController) History(username string) {
	watchlist := new(services.WatchlistService)

	history, err := watchlist.ParseHistory(username)

	if err != nil {
		return
	}

	if len(history.History) > 10 {
		history.History = history.History[:10]
	}

	for _, h := range history.History {
		fmt.Printf("name: %s - episode: %d\n", h.Meta.Name, h.Episode)
	}
}

func (w *WatchlistController) Watchlist(username string) {
	watchlist := new(services.WatchlistService)

	body, err := watchlist.ParseWatchlist(username)

	if err != nil {
		return
	}

	watchlist.ShowWatchlist(body)
}

func (w *WatchlistController) Excel(username string) {
	watchlist := new(services.WatchlistService)

	body, err := watchlist.ParseWatchlist(username)

	if err != nil {
		return
	}

	watchlist.CreateExel(body)
}
