package entity

type PlaylistsRoot struct {
	Playlists Playlists `json:"playlists"`
}

type Playlists struct {
	URL   string     `json:"href"`
	Items []Playlist `json:"items"`
}

type Playlist struct {
	ID            string `json:"id"`
	Collaborative bool   `json:"collaborative"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	URL           string `json:"href"`
}

type PlayListTracks struct {
	Playlist Playlist `json:"playlist"`
	Track    []Track  `json:"tracks"`
}
