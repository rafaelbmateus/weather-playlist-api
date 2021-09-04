package entity

type Tracks struct {
	URL   string      `json:"href"`
	Items []TrackItem `json:"items"`
}

type TrackItem struct {
	AddAt string `json:"added_at"`
	Track Track  `json:"track"`
}

type Track struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"href"`
	PreviewURL string `json:"preview_url"`
}
