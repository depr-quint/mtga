package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// scryfall api: https://scryfall.com/docs/api/cards/search
	total, m20 := get(fmt.Sprintf("https://api.scryfall.com/cards/search?q=set:m20"))

	// do something with it
	_ = total
	_ = m20
}

func get(url string) (total int, cards []card) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var search search
	err = json.Unmarshal(data, &search)
	if err != nil {
		log.Fatalln(err)
	}

	// store them in list/map based on id
	cards = append(cards, search.Cards...) // add cards
	if search.HasMore {
		_, next := get(search.NextPage)
		cards = append(cards, next...)
	}
	return search.TotalCards, cards
}

// cardSearch is the structure of the returned JSON by the card search.
type search struct {
	Object     string `json:"object"` // will be "list"
	TotalCards int    `json:"total_cards"`
	HasMore    bool   `json:"has_more"`
	NextPage   string `json:"next_page"`
	Cards      []card `json:"data"`
}

// card is represents the data of a card returned by the card search.
type card struct {
	Object        string `json:"object"` // will be "card"
	ID            string `json:"id"`
	OracleID      string `json:"oracle_id"`
	MultiverseIDs []int  `json:"multiverse_ids"`
	TCGPlayerID   int    `json:"tcgplayer_id"`
	Name          string `json:"name"`
	Lang          string `json:"lang"`
	ReleasedAt    string `json:"released_at"`
	URI           string `json:"uri"`
	ScryfallURI   string `json:"scryfall_uri"`
	Layout        string `json:"layout"`
	HighResImage  bool   `json:"highres_image"`
	ImageURIs     struct {
		Small      string `json:"small"`
		Normal     string `json:"normal"`
		Large      string `json:"large"`
		PNG        string `json:"png"`
		ArtCrop    string `json:"art_crop"`
		BorderCrop string `json:"border_crop"`
	} `json:"image_uris"`
	ManaCost       string   `json:"mana_cost"`
	CMC            float64  `json:"cmc"`
	TypeLine       string   `json:"type_line"`
	OracleText     string   `json:"oracle_text"`
	Power          string   `json:"power"`
	Toughness      string   `json:"toughness"`
	Colors         []string `json:"colors"`
	ColorIndicator []string `json:"color_indicator"`
	ColorIdentity  []string `json:"color_identity"`
	Legalities     struct {
		Standard  string `json:"standard"`
		Future    string `json:"future"`
		Modern    string `json:"modern"`
		Legacy    string `json:"legacy"`
		Pauper    string `json:"pauper"`
		Vintage   string `json:"vintage"`
		Penny     string `json:"penny"`
		Commander string `json:"commander"`
		Brawl     string `json:"brawl"`
		Duel      string `json:"duel"`
		Oldschool string `json:"oldschool"`
	} `json:"legalities"`
	Games           []string `json:"games"`
	Reserved        bool     `json:"reserved"`
	Foil            bool     `json:"foil"`
	NonFoil         bool     `json:"nonfoil"`
	Oversized       bool     `json:"oversized"`
	Promo           bool     `json:"promo"`
	Reprint         bool     `json:"reprint"`
	Variation       bool     `json:"variation"`
	Set             string   `json:"set"`
	SetName         string   `json:"set_name"`
	SetType         string   `json:"set_type"`
	SetURI          string   `json:"set_uri"`
	SetSearchURI    string   `json:"set_search_uri"`
	ScryfallSetURI  string   `json:"scryfall_set_uri"`
	RulingsURI      string   `json:"rulings_uri"`
	PrintsSearchURI string   `json:"prints_search_uri"`
	CollectorNumber string   `json:"collector_number"`
	Digital         bool     `json:"digital"`
	Rarity          string   `json:"rarity"`
	IllustrationID  string   `json:"illustration_id"`
	CardBackID      string   `json:"card_back_id"`
	Artist          string   `json:"artist"`
	BorderColor     string   `json:"border_color"`
	Frame           string   `json:"frame"`
	FullArt         bool     `json:"full_art"`
	Textless        bool     `json:"textless"`
	Booster         bool     `json:"booster"`
	StorySpotlight  bool     `json:"story_spotlight"`
	Prices          struct {
		USD     string `json:"usd"`
		USDFoil string `json:"usd_foil"`
		EUR     string `json:"eur"`
		TIX     string `json:"tix"`
	} `json:"prices"`
	RelatedURIs struct {
		Gatherer       string `json:"gatherer"`
		TCGPlayerDecks string `json:"tcgplayer_decks"`
		EDHREC         string `json:"edhrec"`
		MTGTop8        string `json:"mtgtop8"`
	} `json:"related_uris"`
	PurchaseURIs struct {
		TCGPlayer   string `json:"tcgplayer"`
		CardMarket  string `json:"cardmarket"`
		CardHoarder string `json:"cardhoarder"`
	} `json:"purchase_uris"`
}
