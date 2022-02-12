package base

type convertSpotifyToAppleRequest struct {
	Id string `json:"id"`
}

type convertSpotifyToAppleResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type convertAppleToSpotifyRequest struct {
	Id string `json:"id"`
}

type convertAppleToSpotifyResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type SpotifyPlaylistResponse struct {
	Href     string   `json:"href"`
	Items    []string `json:"items"`
	Limit    int      `json:"limit"`
	Next     string   `json:"next"`
	Offset   int      `json:"offset"`
	Previous string   `json:"previous"`
	Total    int      `json:"total"`
}

type GetUsersPlaylistsSpotifyResponse struct {
	Href     string   `json:"href"`
	Items    []string `json:"items"`
	Limit    int      `json:"limit"`
	Next     string   `json:"next"`
	Offset   int      `json:"offset"`
	Previous string   `json:"previous"`
	Total    int      `json:"total"`
}
