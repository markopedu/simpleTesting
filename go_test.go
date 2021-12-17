package simpleTesting

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Marko", "")
	want := "Hello Marko"

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestUSGreeting(t *testing.T) {
	got := Greeting("Marko", "en-US")
	want := "Hello Marko"

	if got != want {
		t.Fatalf("US Expected %q, got %q", want, got)
	}
}

func TestFRGreeting(t *testing.T) {
	got := Greeting("Marko", "fr-FR")
	want := "Bonjour Marko"

	if got != want {
		t.Fatalf("FR Expected %q, got %q", want, got)
	}
}

func TestITGreeting(t *testing.T) {
	got := Greeting("Marko", "it-IT")
	want := "Ciao Marko"

	if got != want {
		t.Fatalf("IT Expected %q, got %q", want, got)
	}
}


func TestFrenchGreeting(t *testing.T) {
	got := translate("fr-FR")
	want := "Bonjour "

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}