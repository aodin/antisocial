package network

import (
	"testing"
)

func TestSchema(t *testing.T) {
	if Hoods.Name != "hoods" {
		t.Fatalf("Unexpected name of Hoods table: %s", Hoods.Name)
	}
}
