package monitor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var useragents = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 7.0; Moto G (5) Build/NPPS25.137-93-8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.137 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B554a Safari/9537.53",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:60.0) Gecko/20100101 Firefox/60.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:59.0) Gecko/20100101 Firefox/59.0",
}

// FootLockerCalendar send the response of the calendar
func (m *Monitor) GetFootLockerCalendar(country string, targets []string) ([]FootLockerCalendar, error) {
	client, err := m.getHTTPClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(fmt.Sprintf("https://www.footlocker.%s/INTERSHOP/static/WFS/Footlocker-Site/-/Footlocker/en_US/Release-Calendar/Launchcalendar/launchdata/launchcalendar_feed_all.json", country))

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

	var calendar []FootLockerCalendar
	err = json.Unmarshal(body, &calendar)
	if err != nil {
		return nil, err
	}

	var calendar2 []FootLockerCalendar
	if targets != nil {
		for _, target := range targets {
			for _, v := range calendar {
				if v.Name == target {
					calendar2 = append(calendar2, v)
				}
			}
		}

		return calendar2, nil
	}

	return calendar, nil
}

// FootLockerSize send the response of the calendar
func (m *Monitor) GetFootLockerPrice(link string) (string, error) {
	client, err := m.getHTTPClient()
	if err != nil {
		return "", err
	}

	resp, err := client.Get(link)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := doc.Find("#add-to-cart-form > div > div:nth-child(1) > div.fl-product-details--price > div > span.fl-price--sale > span").Text()

	return s, nil
}
