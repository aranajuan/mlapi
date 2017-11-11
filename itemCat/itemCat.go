package itemCat

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ArtCat functions related to ML's api that get items by categories
type ItemCat struct {
	data string
}

// LoadItems load all items of page one
// Returns 1 on fail
func (a *ItemCat) LoadItems(idCat string) int {
	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA/search?category=" + idCat)
	if err != nil {
		a.data = ""
		return 1
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		a.data = ""
		return 1
	}
	a.data = string(resp_body)
	return 0
}

// GetData returns response of last load
func (a ItemCat) GetData() string {
	return a.data
}

// GetPrices returns an array of float of the prices of the items previously
// loaded
func (a ItemCat) GetPrices() ([]float32, int) {
	result := gjson.Get(a.data, "results.#.price")
	if !result.Exists() || len(result.Array()) == 0 {
		return nil, 1
	}
	res := make([]float32, 0, 5)
	for _, price := range result.Array() {
		priceParsed, err := strconv.ParseFloat(price.String(), 32)
		if err != nil {
			return res, 1
		}
		priceFloat := float32(priceParsed)
		res = append(res, priceFloat)
	}
	return res, 0
}
