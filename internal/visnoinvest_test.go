package internal

import "testing"

func TestConvertStringToFloat64(t *testing.T) {
	v, _ := convertStringToFloat64("R$ 325,95 M")
	t.Fatalf("%f\n", v)
}

func TestVisnoInvestRun(t *testing.T) {
	repo := NewVisnoInvest()

	repo.Run("PETR4")

	t.Fail()
}
