package musicbrainz

type ArtistResult struct {
	Artists []Artist `xml:"artist-list>artist"`
}

type Artist struct {
	Name           string    `xml:"name"`
	ArtistId       string    `xml:"id,attr"`
	Releases       []Release `xml:"release-group-list>release-group"`
	Disambiguation string    `xml:"disambiguation"`
}

type ReleaseResult struct {
	Artist Artist `xml:"artist"`
}

type Release struct {
	Title       string `xml:"title"`
	ReleaseId   string `xml:"id,attr"`
	ReleaseDate string `xml:"first-release-date"`
}
