package musicbrainz

import "testing"

func TestArtistSearch(t *testing.T) {
	result := SearchArtist("Four Tet")

	if result.Artists[0].Name != "Four Tet" {
		t.Errorf("Nope!")
	}

	if len(result.Artists) != 10 {
		t.Errorf("Wrong number of artists found.")
	}
}

func TestGetReleases(t *testing.T) {
	result := GetReleases("3bcff06f-675a-451f-9075-99e8657047e8")

	if len(result.Releases) >= 103 {
		t.Errorf("Not enough releases")
	}
}