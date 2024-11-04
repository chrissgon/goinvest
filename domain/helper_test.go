package domain

import "testing"

func TestDomain_HelperValuePerShare(t *testing.T) {
	have := ValuePerShare(31643000000.00, 13044496930)
	expected := 2.42

	if have != expected {
		t.Fatalf("ValuePerShare should return %v but got %v", expected, have)
	}
}
