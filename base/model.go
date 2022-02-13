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
	Href     string     `json:"href"`
	Items    []Playlist `json:"items"`
	Limit    int        `json:"limit"`
	Next     string     `json:"next"`
	Offset   int        `json:"offset"`
	Previous string     `json:"previous"`
	Total    int        `json:"total"`
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

type SpotifyOAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Playlist struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Followers     Followers    `json:"followers"`
	Href          string       `json:"href"`
	Id            string       `json:"id"`
	Images        []Images     `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotId    string       `json:"snapshot_id"`
	Tracks        []Tracks     `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Images struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Owner struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}

type Tracks struct {
	Href     string `json:"href"`
	Items    []Track  `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type Track struct {
	Album            Album        `json:"album"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      ExternalIds  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	Id               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	Id                   string       `json:"id"`
	Images               Images       `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	AlbumGroup           string       `json:"album_group"`
	Artists              Artists      `json:"artists"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type Artists struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Images       Images       `json:"images"`
	Name         string       `json:"name"`
	Popularity   string       `json:"popularity"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type ExternalIds struct {
	isrc string `json:"isrc"`
	ean  string `json:"ean"`
	upc  string `json:"upc"`
}

type LinkedFrom struct {
}
