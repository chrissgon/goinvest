package fiis

import (
	"reflect"
	"testing"
)

func TestInternal_NewFiis(t *testing.T) {
	have := NewFiis()
	expected := &Fiis{}

	if !reflect.DeepEqual(have, expected) {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}
}
