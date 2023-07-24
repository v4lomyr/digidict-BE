package router

const (
	PublicPath = "/public"

	LanguagePath = "/language"

	DictionaryPath = "/dictionaries/:language"
)

func PublicRoute(path string) string {
	return PublicPath + path
}
