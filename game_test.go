package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test behaviour of small oscillator
func TestRule(t *testing.T) {
	var g game
	g.init(3)

	for x := 0; x < g.size; x++ {
		g.set(x, 1, 1)
	}

	g.run()

	for y := 0; y < g.size; y++ {
		assert.Equal(t, 0, g.get(0, y), "Should be equal")
		assert.Equal(t, 1, g.get(1, y), "Should be equal")
		assert.Equal(t, 0, g.get(2, y), "Should be equal")
	}

	g.run()

	for x := 0; x < g.size; x++ {
		assert.Equal(t, 0, g.get(x, 0), "Should be equal")
		assert.Equal(t, 1, g.get(x, 1), "Should be equal")
		assert.Equal(t, 0, g.get(x, 2), "Should be equal")
	}
}
