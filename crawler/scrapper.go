package crawler

import (
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
	Observations string `json:"observations"`
}

func normalize(code string) string {
	return strings.Trim(code, " ")
}

// Equal -
func (f *FII) Equal(code string) bool {
	return strings.Compare(
		strings.ToLower(f.Code), normalize(strings.ToLower(code))) == 0
}

// Do -
func Do(list chan []FII) {
	Collector := colly.NewCollector(
		colly.CacheDir("./.cache"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	fiis := make([]FII, 0, 150)

	Collector.OnHTML("table tr", func(e *colly.HTMLElement) {
		ch := e.DOM.Children()

		if strings.Contains(ch.Eq(0).Text(), "Código") {
			return
		}

		fii := FII{
			Code:         normalize(ch.Eq(0).Text()),
			BaseDate:     ch.Eq(1).Text(),
			BasePrice:    ch.Eq(2).Text(),
			RealYield:    ch.Eq(4).Text(),
			PaymentDate:  ch.Eq(3).Text(),
			PercentYield: ch.Eq(5).Text(),
			Observations: ch.Eq(6).Text(),
		}

		fiis = append(fiis, fii)
	})

	Collector.Visit(os.Getenv("ENDPOINT"))
	list <- fiis
}
