package main

import (
	"testing"
)

func TestApp_getPricetext(t *testing.T) {

	open, _, _ := testApp.getPriceText()
	if open.Text != "Open: $2587.3500 USD" {
		t.Errorf("expected %s, got %s", "Open: $2587.3500 USD", open.Text)
	}

}
