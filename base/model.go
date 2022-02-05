package base

type convertSpotifyToAppleRequest struct {
	id string `json:"id"`
}

type convertSpotifyToAppleResponse struct {
	status string `json:"status"`
	err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type convertAppleToSpotifyRequest struct {
	id string `json:"id"`
}

type convertAppleToSpotifyResponse struct {
	status string `json:"status"`
	err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type SpotifyPlaylistResponse struct {
	href     string   `json:"href"`
	items    []string `json:"items"`
	limit    int      `json:"limit"`
	next     string   `json:"next"`
	offset   int      `json:"offset"`
	previous string   `json:"previous"`
	total    int      `json:"total"`
}

type GetUsersPlaylistsSpotifyResponse struct {
	href     string   `json:"href"`
	items    []string `json:"items"`
	limit    int      `json:"limit"`
	next     string   `json:"next"`
	offset   int      `json:"offset"`
	previous string   `json:"previous"`
	total    int      `json:"total"`
}
