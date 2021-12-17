package simpleTesting

import "testing"


type GreetingTest struct {
	name string
	locale string
	want string
}

var greetingTests = []GreetingTest {
	{ "George", "en-US", "Hello George" },
	{ "Adam", "fr-FR", "Bonjour Adam" },
	{ "Paolo", "it-IT", "Ciao Paolo" },
	{ "Pekka", "fi-FI", "Terve Pekka" },
}

func TestGreetings(t *testing.T) {

	for _, test := range greetingTests{
		got := Greeting(test.name, test.locale)
		want := test.want

		if got != want{
		t.Fatalf("Expected %q, got %q", want, got)
	}
	}
}

