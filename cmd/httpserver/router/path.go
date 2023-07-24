package router

const (
	PublicPath = "/public"

	LanguagePath = "/language"
)

func PublicRoute(path string) string {
	return PublicPath + path
}
