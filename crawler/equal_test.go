package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_get_true_to_fii_BCFF11(t *testing.T) {
	fii := FII{
		Code: "BCFF11",
	}

	assert.True(t, fii.Equal("BCFF11"))
}

func Test_should_get_true_when_exist_space_in_right_side(t *testing.T) {
	fii := FII{
		Code: "BCFF11",
	}

	assert.True(t, fii.Equal("BCFF11       "))
}

func Test_should_get_true_when_exist_space_in_left_side(t *testing.T) {
	fii := FII{
		Code: "BCFF11",
	}

	assert.True(t, fii.Equal("        BCFF11"))
}

func Test_should_get_false(t *testing.T) {
	fii := FII{
		Code: "BCFF11",
	}

	assert.False(t, fii.Equal("BMLC11B"))
}
