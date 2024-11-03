package vinosinvest

import (
	"reflect"
	"testing"
)

type visnoInvestMock struct{}

func TestNewVisnoInvest(t *testing.T){
	have := NewVisnoInvest()
	expected := &VisnoInvest{}

	if !reflect.DeepEqual(have, expected) {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}
}

