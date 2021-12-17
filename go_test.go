package simpleTesting

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Marko")
	want := "Hello Marko"

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
