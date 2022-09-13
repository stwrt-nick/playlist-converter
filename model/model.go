package model

type convertSpotifyToAppleRequest struct {
	Id           string `json:"id"`
	PlaylistName string `json:"playlistName"`
}

type convertSpotifyToAppleResponse struct {
	Status string `json:"status"`
	Err    error  `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type convertAppleToSpotifyRequest struct {
	Id string `json:"id"`
}

type getAppleJWTTokenRequest struct {
	Id string `json:"id"`
}

type getAppleSongRequest struct {
	SongId string `json:"id"`
}

type getAppleSongResponse struct {
	SongId string `json:"id"`
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

type getAppleJWTTokenResponse struct {
	JWTToken string `json:"jwt_token"`
	Err      error  `json:"err,omitempty"`
}

type ApplePlaylistResponse struct {
	Data []PlaylistResponseData `json:"data"`
}

type PlaylistResponseData struct {
	Id         string                     `json:"id"`
	Type       string                     `json:"type"`
	Href       string                     `json:"href"`
	Attributes PlaylistResponseAttributes `json:"attributes"`
}

type PlaylistResponseAttributes struct {
	HasCatalog  bool        `json:"hasCatalog"`
	Description Description `json:"description"`
	Name        string      `json:"name"`
	CanEdit     bool        `json:"canEdit"`
	IsPublic    bool        `json:"isPublic"`
	PlayParams  PlayParams  `json:"playParams"`
	DateAdded   string      `json:"dateAdded"`
}

type PlayParams struct {
	Id        string `json:"id"`
	Kind      string `json:"kind"`
	IsLibrary bool   `json:"isLibrary"`
}

type Description struct {
	Standard string `json:"standard"`
}

type ApplePlaylistRequest struct {
	Attributes PlaylistAttributes `json:"attributes"`
	Data       []PlaylistData     `json:"data"`
}

type PlaylistData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type PlaylistAttributes struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type GetAppleSongIDByISRCResponse struct {
	Data []Data `json:"data"`
}

type Data struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Href       string     `json:"href"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Previews         []Previews `json:"previews"`
	Artwork          Artwork    `json:"artwork"`
	ArtistName       string     `json:"artistName"`
	Url              string     `json:"url"`
	DiscNumber       string     `json:"discNumber"`
	GenreNames       []string   `json:"genreNames"`
	DurationInMillis int        `json:"durationInMillis"`
	ReleaseDate      string     `json:"releaseDate"`
	Name             string     `json:"name"`
	ISRC             string     `json:"isrc"`
	HasLyrics        bool       `json:"hasLyrics"`
}

type Previews struct {
	Url string `json:"url"`
}

type Artwork struct {
	Width      string `json:"width"`
	Height     string `json:"height"`
	Url        string `json:"url"`
	BgColor    string `json:"bgColor"`
	TextColor1 string `json:"textColor1"`
	TextColor2 string `json:"textColor2"`
	TextColor3 string `json:"textColor3"`
	TextColor4 string `json:"textColor4"`
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
	Tracks        Tracks       `json:"tracks"`
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
	Href     string  `json:"href"`
	Items    []Track `json:"items"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          Artists      `json:"artists"`
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
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewUrl       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	Uri              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
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
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type LinkedFrom struct {
}

type GetPlaylistItems struct {
	Href     string          `json:"href"`
	Items    []PlaylistItems `json:"items"`
	Limit    int             `json:"limit"`
	Next     string          `json:"next"`
	Offset   int             `json:"offset"`
	Previous string          `json:"previous"`
	Total    int             `json:"total"`
}

type PlaylistItems struct {
	AddedAt        string         `json:"added_at"`
	AddedBy        AddedBy        `json:"added_by"`
	IsLocal        bool           `json:"is_local"`
	PrimaryColor   string         `json:"primary_color"`
	Track          Track          `json:"track"`
	VideoThumbnail VideoThumbnail `json:"video_thumbnail"`
}

type AddedBy struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type VideoThumbnail struct {
	Url string `json:"url"`
}
