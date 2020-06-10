package monitor

import (
	"crypto/tls"
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func (m *Monitor) getHTTPClient() (*http.Client, error) {
	rand.Seed(time.Now().Unix())
	var client *http.Client
	if len(m.Proxies) != 0 {
		proxyURL, err := url.Parse("http://" + m.Proxies[rand.Intn(len(m.Proxies))])
		if err != nil {
			return nil, errors.New("Bad proxy URL")
		}
		transport := &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{
			Transport: transport,
			Timeout:   5 * time.Second,
		}
	} else {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
		}
	}

	return client, nil
}
