package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/riquellopes/golf/crawler"
)

// SearchExample -
type SearchExample struct {
	One string `json:"exampleONE"`
	Two string `json:"exampleTWO"`
}

// Home -
type Home struct {
	About         string `json:"about"`
	Contact       string `json:"contact"`
	Project       string `json:"project"`
	SearchExample `json:"search"`
}

// Index -
func Index(c echo.Context) error {
	home := Home{"Recupera informações sobre proventos dos fii",
		"http://www.henriquelopes.com.br",
		"https://github.com/riquellopes/golf",
		SearchExample{"curl -X GET http://vast-lake-49104.herokuapp.com/api/",
			"curl -X GET http://vast-lake-49104.herokuapp.com/api/CNES11B"}}
	return c.JSON(http.StatusOK, home)
}

// CodeF -
func CodeF(list []crawler.FII) echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")

		for _, fii := range list {
			if fii.Equal(code) {
				return c.JSON(http.StatusOK, fii)
			}
		}

		return c.JSON(http.StatusNotFound, "FII Notfound.")
	}
}

// AllF -
func AllF(list []crawler.FII) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, list)
	}
}
