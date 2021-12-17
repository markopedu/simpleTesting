package simpleTesting

func translate(s string) string {
	switch s {
	case "en-US":
		return "Hello "
	case "fr-FR":
		return "Bonjour "
	case "it-IT":
		return "Ciao "
	case "fi-FI":
		return "Terve "
	default:
		return "Hello "
	}
}


func Greeting(s string, locale string) string {
	salutation := translate(locale)
	return salutation + s
}

