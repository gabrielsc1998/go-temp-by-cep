package address

import (
	"reflect"
	"testing"
)

func TestAddressEntity(t *testing.T) {
	code := "12345-678"
	state := "SP"
	city := "Sao Paulo"
	district := "Centro"
	addr := "Rua XYZ"

	expected := &Address{
		Code:     code,
		State:    state,
		City:     city,
		District: district,
		Address:  addr,
	}

	result := New(code, state, city, district, addr)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
