package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func (m *Monitor) GetShopifyProducts(target string) (*ShopifyProducts, error) {
	client, err := m.getHTTPClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(fmt.Sprintf("%s/products.json", target))

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

	var products ShopifyProducts
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}
