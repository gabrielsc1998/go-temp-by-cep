package temp

import "testing"

func TestTemperatureEntity(t *testing.T) {
	temp := New(10, 50)
	if temp.TempC != 10 {
		t.Errorf("expected 10, got %f", temp.TempC)
	}
	if temp.TempF != 50 {
		t.Errorf("expected 50, got %f", temp.TempF)
	}
	if temp.TempK != 283.15 {
		t.Errorf("expected 283.15, got %f", temp.TempK)
	}
}
