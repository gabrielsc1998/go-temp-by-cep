package cep_validators

import (
	"testing"
)

func TestIsValidCep(t *testing.T) {
	valid := IsValidCep("12345-678")
	if !valid {
		t.Errorf("Expected true, got '%t'", valid)
	}

	valid = IsValidCep("12345678")
	if !valid {
		t.Errorf("Expected true, got '%t'", valid)
	}

	valid = IsValidCep("abcd")
	if valid {
		t.Errorf("Expected false, got '%t'", valid)
	}
}
