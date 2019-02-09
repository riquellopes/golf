package api

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/riquellopes/golf/crawler"
	"github.com/stretchr/testify/assert"
)

var (
	fiiJSON = `[{"code":"BCFF11","base_date":"","base_price":"","real_yield":"","payment_date":"","percent_yield":"","observations":""}]`
)

func Test_should_get_status_200(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/api/BCFF11", nil)
	rec := test.NewResponseRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/api/:code")
	context.SetParamNames("code")
	context.SetParamValues("BCFF11")

	fii := []crawler.FII{
		crawler.FII{
			Code: "BCFF11",
		},
	}

	assert.NoError(t, CodeF(fii)(context))
	assert.Equal(t, http.StatusOK, rec.Status())
}

func Test_should_get_status_404(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	context := e.NewContext(req, rec)
	context.SetPath("/api/:code")
	context.SetParamNames("code")
	context.SetParamValues("BCFF112")

	fii := []crawler.FII{
		crawler.FII{
			Code: "BCFF11",
		},
	}

	assert.NoError(t, CodeF(fii)(context))
	assert.Equal(t, http.StatusNotFound, rec.Status())
}

func Test_should_get_a_list_of_items(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	context := e.NewContext(req, rec)

	fii := []crawler.FII{
		crawler.FII{
			Code: "BCFF11",
		},
	}

	assert.NoError(t, AllF(fii)(context))
	assert.Equal(t, http.StatusOK, rec.Status())
	assert.Equal(t, fiiJSON, rec.Body.String())
}

func Test_should_get_200_when_to_call_the_index(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	context := e.NewContext(req, rec)

	assert.NoError(t, Index(context))
	assert.Equal(t, http.StatusOK, rec.Status())
}
