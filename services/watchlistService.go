package services

import (
	"encoding/json"
	"errors"
	"exporter/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

type WatchlistService struct {
}

var httpClient = &http.Client{
	Timeout: time.Second * 15,
}

func (s *WatchlistService) ParseHistory(username string) (models.History, error) {
	history := models.History{}

	url := fmt.Sprintf("https://api.jikan.moe/v3/user/%s/history/anime", username)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")

	if err != nil {
		return history, err
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return history, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New("user is not found")
		log.Fatal(err)
		return history, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(resp.StatusCode)
		return history, err
	}

	json.Unmarshal(body, &history)

	return history, nil
}

func (w *WatchlistService) CreateJson(body []byte) error {
	f, err := os.Create("watchlist.json")

	if err != nil {
		return err
	}

	f.Write(body)

	if err != nil {
		return err
	}

	return nil
}

func (w *WatchlistService) ParseWatchlist(username string) ([]byte, error) {
	if len(username) <= 0 {
		err := errors.New("username length must not be 0")
		log.Fatal(err)
		return make([]byte, 0), err
	}

	url := fmt.Sprintf("https://api.jikan.moe/v3/user/%s/animelist/all", username)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")

	if err != nil {
		return make([]byte, 0), err
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New("user is not found")
		log.Fatal(err)
		return make([]byte, 0), err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(resp.StatusCode)
		return make([]byte, 0), err
	}

	return body, nil
}

func (w *WatchlistService) CreateExel(body []byte) {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "mal_id")
	f.SetCellValue("Sheet1", "B1", "name")
	f.SetCellValue("Sheet1", "C1", "score")
	f.SetCellValue("Sheet1", "D1", "url")
	f.SetCellValue("Sheet1", "E1", "watched_episodes")
	f.SetCellValue("Sheet1", "F1", "total_episodes")
	f.SetCellValue("Sheet1", "G1", "type")

	anime := models.WatchlistExel{}
	cell := 1

	json.Unmarshal(body, &anime)

	for _, a := range anime.Anime {
		cell = cell + 1
		cellNumber := fmt.Sprintf("%d", cell)
		f.SetCellValue("Sheet1", "A"+cellNumber, a.ID)
		f.SetCellValue("Sheet1", "B"+cellNumber, a.Name)
		f.SetCellValue("Sheet1", "C"+cellNumber, a.Score)
		f.SetCellValue("Sheet1", "D"+cellNumber, a.Url)
		f.SetCellValue("Sheet1", "E"+cellNumber, a.Wepiosedes)
		f.SetCellValue("Sheet1", "F"+cellNumber, a.Tepisodes)
		f.SetCellValue("Sheet1", "G"+cellNumber, a.Type)
	}

	if err := f.SaveAs("watchlist.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func (w *WatchlistService) ShowWatchlist(body []byte) {
	anime := models.WatchListJson{}

	json.Unmarshal(body, &anime)

	for _, a := range anime.Anime {
		fmt.Printf("name: %s - watched_episodes: %d total_episodes: %d\n", a.Name, a.Wepisodes, a.Aepisodes)
	}
}