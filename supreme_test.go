package monitor

import (
	"log"
	"testing"
)

func TestMonitor_GetSupremeProducts(t *testing.T) {
	monitor := NewMonitor("discord_webhook", nil)

	productLinks, err := monitor.GetSupremeProducts()

	if err != nil {
		log.Fatal(err)
	}
	for _, product := range productLinks.ProductsAndCategories.Bags {
		log.Print(product.ID)
	}
}

func TestMonitor_GetSupremeProduct(t *testing.T) {
	monitor := NewMonitor("discord_webhook", nil)

	productInfo, err := monitor.GetSupremeProduct(305051)

	if err != nil {
		log.Fatal(err)
	}

	for _, style := range productInfo.Styles {
		for _, size := range style.Sizes {
			log.Printf("Stock for %s: %v", size.Name, size.StockLevel)
		}
	}
}

func TestMonitor_GetSupremeExample(t *testing.T) {
	monitor := NewMonitor("discord_webhook", nil)

	productLinks, err := monitor.GetSupremeProducts()

	if err != nil {
		log.Fatal(err)
	}
	for _, products := range productLinks.ProductsAndCategories.Jackets {
		productInfo, err := monitor.GetSupremeProduct(products.ID)
		if err != nil {
			log.Fatal(err)
		}

		for _, style := range productInfo.Styles {
			for _, size := range style.Sizes {
				log.Printf("Stock for %s %s (%s): %v. Blocked to EU: %v, RU: %v, CA: %v. Price: %d$ / %d€", products.Name, style.Name, size.Name, size.StockLevel, productInfo.NonEuBlocked, productInfo.RussiaBlocked, productInfo.CanadaBlocked, products.Price/100, products.PriceEuro/100)
			}
		}
	}
}
