package musicbrainz

type ArtistResult struct {
	Artists []Artist `xml:"artist-list>artist"`
}

type Artist struct {
	Name           string `xml:"name"`
	Id             string `xml:"id,attr"`
	Disambiguation string `xml:"disambiguation"`
}

type ReleaseResult struct {
	Releases []Release `xml:"release-list>release"`
}

type Release struct {
	Title string `xml:"title"`
	Id    string `xml:"id,attr"`
}
