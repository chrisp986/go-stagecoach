package service

import "testing"

func TestValidateEmail(t *testing.T) {
	got := validateEmail("test@test.com")
	want := true

	if got != want {
		t.Error("Validation if argument is mail address")
	}
}

func TestValidateNotEmpty(t *testing.T) {
	got := validateNotEmpty("teststring")
	want := true

	if got != want {
		t.Error("Valdiation if argument is not empty")
	}
}
