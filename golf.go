package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// SearchExemple -
type SearchExemple struct {
	One string `json:"exampleONE"`
	Two string `json:"exampleTWO"`
}

// Home -
type Home struct {
	About         string `json:"about"`
	Contact       string `json:"contact"`
	Project       string `json:"project"`
	SearchExemple `json:"search"`
}

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		home := Home{"Recupera informações sobre proventos dos fii",
			"http://www.henriquelopes.com.br",
			"https://github.com/riquellopes/fii",
			SearchExemple{"curl -X GET http://vast-lake-49104.herokuapp.com/api/",
				"curl -X GET http://vast-lake-49104.herokuapp.com/api/CNES11B"}}
		return c.JSON(http.StatusOK, home)
	})

	app.GET("/api/", allFII("list alls"))
	app.GET("/api/:codigo", fiiByCode("one"))

	app.Logger.Fatal(app.Start(":5000"))
}

func allFII(name string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, name)
	}
}

func fiiByCode(name string) echo.HandlerFunc {
	return func(c echo.Context) error {
		codigo := c.Param("codigo")

		return c.String(http.StatusOK, codigo)
	}
}
