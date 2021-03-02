package main

import (
	"testing"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func FakeLoad (a MagicAPI) {
	a.Awesome("bob")
}

func TestHello(t *testing.T) {
	// empty -- confirm compilation only
	m := new(MagicAPI)
	m.On("Awesome", "bob").Return("what about")

	FakeLoad(m)

	m.AssertExpectations(t)
	//assert.Called(
}
