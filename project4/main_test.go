package main

import (
	_ "log"
	"testing"
)

func TestHighestPalidrone(t *testing.T) {
	hp := highestPalidrone(99, 99)

	if hp != 9009 {
		t.Fail()
	}
}
