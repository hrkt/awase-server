package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	event := NewEvent("myevent")
	assert.Equal(t, "myevent", event.Label)
	blank := []CandidateDate{}
	assert.Equal(t, blank, event.CandidateDates)
}
