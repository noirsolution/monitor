package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)



func (m *Monitor) GetSupremeProducts() (*SupremeProducts, error) {
	client, err := m.getHTTPClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get("https://www.supremenewyork.com/mobile_stock.json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("invalid response code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var products SupremeProducts
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (m *Monitor) GetSupremeProduct(id int) (*SupremeProduct, error) {
	client, err := m.getHTTPClient()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.supremenewyork.com/shop/%v.json", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authority", "www.supremenewyork.com")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("invalid response code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var product SupremeProduct
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
