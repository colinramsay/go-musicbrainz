package musicbrainz

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func MakeQuery(url string) []byte {
	log.Printf("Getting URL %s", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func SearchArtist(artist string) ArtistResult {
	result := ArtistResult{}

	bytes := MakeQuery("http://musicbrainz.org/ws/2/artist/?limit=10&query=artist:" + url.QueryEscape(artist))

	xml.Unmarshal(bytes, &result)

	return result
}

func GetReleases(artistId string) ReleaseResult {
	result := ReleaseResult{}
	bytes := MakeQuery("http://musicbrainz.org/ws/2/release?artist=" + artistId)

	xml.Unmarshal(bytes, result)

	return result
}
