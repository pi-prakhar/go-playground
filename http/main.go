package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
)

const BASE_URL = "https://deckofcardsapi.com/api/deck/"

type Deck struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type Card struct {
	Code   string `json:"code"`
	Image  string `json:"image"`
	Images struct {
		Svg string `json:"svg"`
		Png string `json:"png"`
	} `json:"images"`
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

type DrawResponse struct {
	Success   bool   `json:"success"`
	DeckID    string `json:"deck_id"`
	Cards     []Card `json:"cards"`
	Remaining int    `json:"remaining"`
}

var priorityMap map[string]int = map[string]int{
	"A": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"0": 10, // 10 is represented as "0"
	"J": 11,
	"Q": 12,
	"K": 13,
}

func getNewDeck() (Deck, error) {
	url := BASE_URL + "/new"
	var newDeck Deck
	resp, err := http.Get(url)

	if err != nil {
		return newDeck, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&newDeck)
	if err != nil {
		return newDeck, err
	}

	return newDeck, nil
}

func drawCards(deckID string, count int) ([]Card, error) {
	var cards []Card
	var drawResponse DrawResponse
	url := fmt.Sprintf("%s%s%s%s%d", BASE_URL, deckID, "/draw/", "?count=", count)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return cards, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&drawResponse)
	if err != nil {
		return cards, err
	}

	return drawResponse.Cards, nil
}

func drawCardsNew(deckID string, count string) ([]Card, error) {
	var cards []Card
	var drawResponse DrawResponse

	endpoint := fmt.Sprintf("%s/%s/draw", BASE_URL, deckID)
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal("Error parsing URL:", err)
		return cards, err
	}

	params := url.Values{}
	params.Add("count", count)

	u.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return cards, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cards, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&drawResponse)
	if err != nil {
		return cards, err
	}

	return drawResponse.Cards, nil

}

func drawCardsFromNewDeck(count int) ([]Card, error) {
	var cards []Card
	var drawResponse DrawResponse
	url := fmt.Sprintf("%s%s%s%s%d", BASE_URL, "new/", "draw/", "?count=", count)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return cards, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&drawResponse)
	if err != nil {
		return cards, err
	}

	return drawResponse.Cards, nil
}

func shuffle(deckID string) error {
	var deck Deck
	endpoint := BASE_URL + deckID + "/shuffle/"
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&deck)
	if err != nil {
		return err
	}

	return nil
}

func returnAllCards(deckID string) error {
	var deck Deck
	endpoint := BASE_URL + deckID + "/return/"
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&deck)
	if err != nil {
		return err
	}

	fmt.Println(deck)

	return nil
}

func main() {
	deck, err := getNewDeck()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(deck)

	err = shuffle(deck.DeckID)
	if err != nil {
		log.Println(err.Error())
	}

	cards, err := drawCardsNew(deck.DeckID, "6")
	if err != nil {
		log.Println(err.Error())
	}

	for _, card := range cards {
		fmt.Println(card.Code)
	}

	sort.Slice(cards, func(i, j int) bool {
		return priorityMap[cards[i].Code[:1]] > priorityMap[cards[j].Code[:1]]
	})

	for _, card := range cards {
		fmt.Println(card.Code)
	}
	// cards, err = drawCardsFromNewDeck(2)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(cards)

	// err = returnAllCards(deck.DeckID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
