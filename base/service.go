package base

type service struct{}

type Service interface {
	ConvertSpotifyToApple(string) string
	ConvertAppleToSpotify(string) string
}

func ConvertSpotifyToApple(string) string {
	return "Hello"
}

func ConvertAppleToSpotify(string) string {
	return "Hello"
}
