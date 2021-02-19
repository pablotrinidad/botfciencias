package lib

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL        = "https://pagina.fciencias.unam.mx/"
	httpUserAgent  = "BotFCiencias v.0.1 https://github.com/pablotrinidad/botfciencias"
	requestTimeout = time.Duration(2 * time.Second)
)

func loadPageContent(path string, out interface{}) error {
	client := getHTTPClient()
	request, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", httpUserAgent)

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data string

	doc.Find("head script").Each(func(i int, selection *goquery.Selection) {
		if len(selection.Nodes) != 1 {
			return
		}
		node := selection.Nodes[0]
		selection = selection.FilterNodes(node)
		if _, ok := selection.Attr("data-drupal-selector"); !ok {
			return
		}
		data = node.FirstChild.Data
	})

	if data == "" {
		return fmt.Errorf("failed finding content on %s", path)
	}

	return json.Unmarshal([]byte(data), out)
}

func getHTTPClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &http.Client{
		Transport: transport,
		Timeout:   requestTimeout,
	}
}
