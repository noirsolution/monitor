package monitor

import (
	"log"
	"testing"
)

func TestMonitor_GetShopifyProducts(t *testing.T) {
	monitor := NewMonitor("x", nil)

	products, err := monitor.GetShopifyProducts("https://testmonitorkaa.myshopify.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(products)
}
