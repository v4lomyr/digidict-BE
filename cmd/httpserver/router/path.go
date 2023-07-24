package router

const (
	PublicPath = "/public"

	LanguagePath = "/language"

	DictionaryPath = "/dictionaries/:language"

	TranslatePath = "/translate"
)

func PublicRoute(path string) string {
	return PublicPath + path
}
