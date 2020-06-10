package monitor

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestFootLocker(t *testing.T) {
	monitor := NewMonitor("webhook_url", nil)
	var inStock bool
	var target FootLockerCalendar
	for !inStock {
		monitorTarget := []string{"Kyrie 6"} // In Stock
		calendar, err := monitor.GetFootLockerCalendar("de", monitorTarget)
		if err != nil {
			log.Fatal(err)
		}
		target = calendar[0]

		inStock, err = strconv.ParseBool(target.HasStock)
		if err != nil {
			log.Fatal(err)
		}
	}

	var price string

	for _, country := range target.DeepLinks {
		if country.Locale == "de" {
			var err error
			price, err = monitor.GetFootLockerPrice(country.Link)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	var embedCountry string

	for _, country := range target.DeepLinks {
		embedCountry += fmt.Sprintf("[%s](%s)\n", strings.ToUpper(country.Locale), country.Link)
	}

	err := monitor.SendDiscordWebhook(Webhook{
		Username: "VoidAIO",
		Embeds: []Embeds{
			{
				Title:     target.Name,
				Thumbnail: Thumbnail{URL: target.Image},
				Fields: []Fields{
					{
						Name:   "Country",
						Value:  embedCountry,
						Inline: true,
					},
					{
						Name:   "In Stock",
						Value:  target.HasStock,
						Inline: true,
					},
					{
						Name:   "Color",
						Value:  target.Colors,
						Inline: true,
					},
					{
						Name:   "Price",
						Value:  price,
						Inline: true,
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
