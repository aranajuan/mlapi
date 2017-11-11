package itemCat

import (
	"net/http"
	"net/url"
	"testing"
)

func TestLoadItems(t *testing.T) {
	var ic ItemCat
	ic.LoadItems("MLA5726")
	if ic.GetData() == "" {
		t.Error("LoadItems not getting data")
	}
}

func TestLoadConnectionError(t *testing.T) {
	var ic ItemCat
	proxyUrl, _ := url.Parse("http://proxyIp:proxyPort")
	tmp := http.DefaultTransport
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	if ic.LoadItems("MLA5726") != 1 {
		t.Error("Not error thrown on connection fail")
	}
	http.DefaultTransport = tmp
}

func TestGetPrice(t *testing.T) {
	var ic ItemCat
	ic.LoadItems("MLA5726")
	prices, err := ic.GetPrices()
	if err != 0 || len(prices) == 0 {
		t.Error("Error getting prices")
	}
}

func TestGetPriceNoLoad(t *testing.T) {
	var ic ItemCat
	_, err := ic.GetPrices()
	if err != 1 {
		t.Error("Not error thrown on empty response")
	}
}

func BenchmarkLoad(b *testing.B) {
	var ic ItemCat
	b.N = 2
	for i := 0; i < b.N; i++ {
		ic.LoadItems("MLA5726")
	}
}

func BenchmarkGetPrice(b *testing.B) {
	var ic ItemCat
	ic.LoadItems("MLA5726")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ic.GetPrices()
	}
}
