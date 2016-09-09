package main

import (
   "net/http"
   "github.com/labstack/echo"
   "github.com/labstack/echo/engine/standard"
)

type SearchExemple struct {
    One string `json:"exampleONE"`
    Two string `json:"exampleTWO"`
}

type Home struct {
    About string `json:"about"`
    Contact string `json:"contact"`
    Project string `json:"project"`
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

    app.GET("/scrap", func(c echo.Context) error {
        return c.String(http.StatusOK, "scrap")
    })

    app.GET("/:codigo", func(c echo.Context) error {
        codigo := c.Param("codigo")

        return c.String(http.StatusOK, codigo)
    })

    app.Run(standard.New(":5000"))
}
