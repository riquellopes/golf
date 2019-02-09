package api

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/riquellopes/golf/crawler"
	"github.com/stretchr/testify/assert"
)

func Test_should_not_get_error(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.GET, "/api/BCFF11", nil)
	rec := test.NewResponseRecorder()
	context := e.NewContext(req, rec)

	fii := []crawler.FII{
		crawler.FII{
			Code: "BCFF11",
		},
	}

	assert.NoError(t, CodeF(fii)(context))
}
