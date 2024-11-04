package vinosinvest

import (
	"reflect"
	"testing"
)

func TestInternal_NewVisnoInvest(t *testing.T){
	have := NewVisnoInvest()
	expected := &VisnoInvest{}

	if !reflect.DeepEqual(have, expected) {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}
}

