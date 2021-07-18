package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

type SpringerCreator struct {
	Name string `json:"creator"`
}

type SpringerAbstract struct {
	RawDisplayName json.RawMessage `json:"p"`
}

type SpringerUrl struct {
	Value string `json:"value"`
}

type SpringerRecords struct {
	Title         string            `json:"title"`
	Id            string            `json:"identifier"`
	Creators      []SpringerCreator `json:"creators"`
	PublishedDate string            `json:"publicationDate"`
	Publisher     string            `json:"publisher"`
	Abstract      SpringerAbstract  `json:"abstract"`
	URLs          []SpringerUrl     `json:"url"`
}

type SpringerFeed struct {
	Records []SpringerRecords `json:"records"`
}

/*
   Get a set of Springer articles
   http://api.springernature.com/openaccess/json?api_key=66d2a126a14009044fb70d5781ebb284&q=Carbon Nanotube
   q => Search query
   p => Number of results to return
   s => The start index of results
*/
func GetSpringerArticles(query_string string, amount int, page int) []*pb.Result {
	start := (amount * (page - 1)) + 1
	api_key := os.Getenv("SPRINGER_API_KEY")
	url_template := "http://api.springernature.com/openaccess/json?q=%[1]v&p=%[2]v&s=%[3]v&api_key=%[4]v"
	url := fmt.Sprintf(url_template, query_string, amount, start, api_key)
	fmt.Printf("Making request to: %v \n", url)

	result := make([]*pb.Result, 0)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status error: %v \n", resp.StatusCode)
		return result
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read body error: %v \n", err)
		return result
	}

	// Unmarshal returned XML
	var dataAsJson SpringerFeed
	json.Unmarshal(data, &dataAsJson)

	for _, article := range dataAsJson.Records {
		authors := ""
		for i, creator := range article.Creators {
			if i < len(article.Creators) {
				authors = authors + creator.Name + ";"
			} else {
				authors = authors + creator.Name
			}
		}

		// Handle the case that the p tag is either a string OR array
		description := ""
		rawDisplayName := article.Abstract.RawDisplayName
		if len(rawDisplayName) > 0 {
			switch rawDisplayName[0] {
			case '"':
				if err := json.Unmarshal(rawDisplayName, &description); err != nil {
					log.Fatalf("Error unmarshalling single p tag: %v \n", err)
				}
			case '[':
				var s []string
				if err := json.Unmarshal(rawDisplayName, &s); err != nil {
					log.Fatalf("Error unmarshalling array p tag: %v \n", err)
				}
				// Join arrays with "&" per OP's comment on the question.
				description = strings.Join(s, "\n")
			}
		}

		newArticle := pb.Result{
			Id:            article.Id,
			Author:        authors,
			Title:         article.Title,
			PublishedDate: article.PublishedDate,
			Publisher:     article.Publisher,
			Description:   description,
			Url:           article.URLs[0].Value,
		}
		result = append(result, &newArticle)
	}

	return result
}
