package main

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPrice(t *testing.T) {
	go startServer(false)
	resp, _ := http.Get("http://localhost:8080/categories/MLA5726/price")
	resp_body, _ := ioutil.ReadAll(resp.Body)
	data := string(resp_body)
	max, min, sg := gjson.Get(data, "max"),
		gjson.Get(data, "min"),
		gjson.Get(data, "suggested")
	if !max.Exists() || !min.Exists() || !sg.Exists() {
		t.Error("Not getting correct data")
	}
}

func TestGetPriceNoItems(t *testing.T) {
	go startServer(false)
	resp, _ := http.Get("http://localhost:8080/categories/sfg4tw/price")
	if resp.StatusCode != 404 {
		t.Error("Not error thrown on bad category", resp.StatusCode)
	}

}

func BenchmarkPrice(b *testing.B) {
	b.N = 3
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8080/categories/MLA5726/price")
	}
}
