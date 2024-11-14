package internal

import "testing"

func TestInternal_HelperConvertStringToFloat64(t *testing.T) {
	have, _ := ConvertStringToFloat64("R$ 32,30")
	expected := 32.30

	if have != expected {
		t.Fatalf("ConvertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = ConvertStringToFloat64("R$ 79,01 K")
	expected = 79010.00

	if have != expected {
		t.Fatalf("ConvertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = ConvertStringToFloat64("R$ 325,95 M")
	expected = 325950000.00

	if have != expected {
		t.Fatalf("ConvertStringToFloat64 should return %v but got %v", expected, have)
	}

	have, _ = ConvertStringToFloat64("R$ 2,67 B")
	expected = 2670000000.00

	if have != expected {
		t.Fatalf("ConvertStringToFloat64 should return %v but got %v", expected, have)
	}

	_, err := ConvertStringToFloat64("")

	if err == nil {
		t.Fatalf("ConvertStringToFloat64 should return an error because the param is invalid")
	}
}

func TestInternal_Normalization(t *testing.T) {
	have := Normalization("áéíóúãõâêîôû ÁÉÍÓÚÃÕÂÊÎÔÛ")
	expected := "aeiouaoaeiou AEIOUAOAEIOU"

	if have != expected {
		t.Fatalf("Normalization should return %v but got %v", expected, have)
	}
}

func TestInternal_ConvertTaxStringToFloat64(t *testing.T) {
	have, _ := ConvertTaxStringToFloat64("0.6")
	expected := 0.6

	if have != expected {
		t.Fatalf("ConvertTaxStringToFloat64 should return %v but got %v", expected, have)
	}
	
	have, _ = ConvertTaxStringToFloat64("1.2")
	expected = 1.2

	if have != expected {
		t.Fatalf("ConvertTaxStringToFloat64 should return %v but got %v", expected, have)
	}
	
	have, _ = ConvertTaxStringToFloat64("0,9")
	expected = 0.9

	if have != expected {
		t.Fatalf("ConvertTaxStringToFloat64 should return %v but got %v", expected, have)
	}
}