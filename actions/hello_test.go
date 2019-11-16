package actions

import (
	"testing"
)

func TestGreet(t *testing.T) {
	str := Greet()
	if str != "smoke"{
		t.Errorf("Greet() => %s, Expected hello github action", str)
	}
}
