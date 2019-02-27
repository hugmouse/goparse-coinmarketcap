package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/logrusorgru/aurora"
)

// Coin is a coin obviously
// type Coin struct {
// ID     string
// Name   string
// Symbol string
// CmcRank           int
// NumMarketPairs    int
// CirculatingSupply int
// TotalSupply       int
// MaxSupply         int
// LastUpdated       string
// DateAdded         string
// Tags              []string
// Platform          string
// }

func main() {
	fName := "cryptocoinmarketcap.csv"
	fNameJSON := "cryptocoinmarketcap.json"
	var IDK uint
	var A []string
	//var AJSON []Coin
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	fileJSON, err := os.Create(fNameJSON)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fNameJSON, err)
		return
	}
	defer fileJSON.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"ID", "Name", "Symbol", "Market capacity (USD)", "Market capacity (BTC)", "Price (USD)", "Price (BTC)", "Circulating Supply (USD)", "Volume (USD)", "Volume (BTC)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnXML("//tbody/tr", func(e *colly.XMLElement) {
		A = []string{
			e.ChildText("/td[@class='text-center']"),                                        // ID
			e.ChildText("//a[@class='currency-name-container link-secondary']"),             // Name
			e.ChildText("/td[@class='text-left col-symbol']"),                               // Symbol
			e.ChildAttr("/td[@class='no-wrap market-cap text-right']", "data-usd"),          // Market Capacity (USD)
			e.ChildAttr("/td[@class='no-wrap market-cap text-right']", "data-btc"),          // Market Capacity (BTC)
			e.ChildAttr("/td[@class='no-wrap text-right']/a", "data-usd"),                   // Price (USD)
			e.ChildAttr("/td[@class='no-wrap text-right']/a", "data-btc"),                   // Price (BTC)
			e.ChildAttr("/td[@class='no-wrap text-right circulating-supply']", "data-sort"), // Cicrulating Supply (USD)
			e.ChildAttr("/td[@class='no-wrap text-right ']/a", "data-usd"),                  // Volume (USD)
			e.ChildAttr("/td[@class='no-wrap text-right ']/a", "data-btc"),                  // Volume (BTC)
			e.ChildAttr("/td[@data-timespan='1h']", "data-sort"),                            // Change (1h)
			e.ChildAttr("/td[@data-timespan='24h']", "data-sort"),                           // Change (24h)
			e.ChildAttr("/td[@data-timespan='7d']", "data-sort"),                            // Change (7d)
		}

		IDK++
		// TODO: add multi-dimension JSON output.
		// AJSON = append(AJSON, Coin{
		// 	ID:     A[0],
		// 	Name:   A[1],
		// 	Symbol: A[2],
		// })
		json.NewEncoder(fileJSON).Encode(A)
		writer.Write(A)
		// fmt.Printf("\rScraping %d", IDK)
		fmt.Printf("\rAmount of cryptocurrencies: %d", aurora.Green(IDK))
	})
	//json.NewEncoder(io.Writer(fileJSON)).Encode(AJSON)
	c.Visit("https://coinmarketcap.com/all/views/all/")
	fmt.Printf("\nScraping finished, check file %q and %q for results\n", aurora.Green(fName), aurora.Green(fNameJSON))
}
