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

func TestConvertStringToFloat64(t *testing.T) {
	have, _ := convertStringToFloat64("R$ 32,30")
	expected := 32.30

	if have != expected {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = convertStringToFloat64("R$ 79,01 K")
	expected = 79010.00

	if have != expected {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = convertStringToFloat64("R$ 325,95 M")
	expected = 325950000.00

	if have != expected {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = convertStringToFloat64("R$ 2,67 B")
	expected = 2670000000.00

	if have != expected {
		t.Fatalf("convertStringToFloat64 should return %v but got %v", expected, have)
	}

	_, err := convertStringToFloat64("")

	if err == nil {
		t.Fatalf("convertStringToFloat64 should return an error because the param is invalid")
	}
}
