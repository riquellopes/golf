package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCollector struct {
}

func (c *MockCollector) Extract() []FII {
	fiis := make([]FII, 0)

	return append(fiis, FII{Code: "BCFF11"})
}

func Test_shoulg_get_equal_one(t *testing.T) {

	fiis := make(chan []FII)
	defer close(fiis)

	go Do(fiis, new(MockCollector))
	output := <-fiis

	assert.Equal(t, len(output), 1)
}
