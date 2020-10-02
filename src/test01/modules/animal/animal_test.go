package animal

import (
	"testing"
)

func TestDogBark(t *testing.T) {
	want := "멍멍"
	if got := DogBark(); got != want {
		t.Errorf("Hello()=%q, want %q", got, want)
	}
}

func TestCatMeow(t *testing.T) {
	want := "냐옹냐옹"
	if got := CatMeow(); got != want {
		t.Errorf("CatMeow()=%q, want %q", got, want)
	}
}
