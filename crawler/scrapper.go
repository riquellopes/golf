package crawler

import (
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// FII -
type FII struct {
	Code         string `json:"code"`
	BaseDate     string `json:"base_date"`
	BasePrice    string `json:"base_price"`
	RealYield    string `json:"real_yield"`
	PaymentDate  string `json:"payment_date"`
	PercentYield string `json:"percent_yield"`
}

func normalize(code string) string {
	return strings.Trim(code, " ")
}

func getCode(url string) string {
	return strings.ToUpper(strings.Replace(url, "/", "", 2))
}

// Equal -
func (f *FII) Equal(code string) bool {
	return strings.Compare(
		strings.ToLower(f.Code), normalize(strings.ToLower(code))) == 0
}

// Collector -
type Collector interface {
	Extract() []FII
}

// FiiCollector -
type FiiCollector struct {
}

// Extract -
func (f *FiiCollector) Extract() []FII {
	Collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.Async(true),
		colly.MaxDepth(2),
	)

	fiis := make([]FII, 0)

	Collector.OnHTML("div#items-wrapper div", func(e *colly.HTMLElement) {
		ch := e.DOM.Children()
		link, _ := ch.Attr("href")

		// fmt.Println(link)

		Collector.Visit(e.Request.AbsoluteURL(link))
	})

	Collector.OnHTML("#last-revenues--table > tbody > tr:nth-child(1)", func(e *colly.HTMLElement) {
		ch := e.DOM.Children()
		code := getCode(e.Request.URL.Path)

		fii := FII{
			Code:         code,
			BaseDate:     ch.Eq(0).Text(),
			BasePrice:    ch.Eq(2).Text(),
			RealYield:    ch.Eq(3).Text(),
			PaymentDate:  ch.Eq(1).Text(),
			PercentYield: ch.Eq(4).Text(),
		}

		fiis = append(fiis, fii)
	})

	Collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	Collector.Visit(os.Getenv("ENDPOINT") + "/lista-de-fundos-imobiliarios/")
	return fiis
}

// Do -
func Do(list chan []FII, collector Collector) {
	list <- collector.Extract()
}
