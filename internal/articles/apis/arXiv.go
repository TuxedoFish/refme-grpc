package apis

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "github.com/TuxedoFish/refme-grpc/pkg/proto/articlespb"
)

type ArxivURL struct {
	Link        string `xml:"href,attr"`
	ContentType string `xml:"type,attr"`
}

type ArxivArticle struct {
	Id            string     `xml:"id"`
	Title         string     `xml:"title"`
	PublishedDate string     `xml:"published"`
	Description   string     `xml:"summary"`
	URLs          []ArxivURL `xml:"link"`
	Authors       []string   `xml:"author>name"`
}

type ArxivFeed struct {
	Articles []ArxivArticle `xml:"entry"`
}

/*
   Get a set of arxiv articles into a standard format
   Example: http://export.arxiv.org/api/query?search_query=all:carbon%20nanotube&start=0&max_results=5
*/
func GetArXivArticles(query_string string, amount int, page int) []*pb.Result {
	url_template := "https://export.arxiv.org/api/query?search_query=%[1]v&max_results=%[2]v&start=%[3]v"
	url := fmt.Sprintf(url_template, query_string, amount, (page-1)*amount)
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
	var dataAsXML ArxivFeed
	xml.Unmarshal(data, &dataAsXML)

	// Loop over putting them into Result objects
	for _, article := range dataAsXML.Articles {
		// Fetch link pointing to PDF
		articleURL := "unknown"
		for _, linkTag := range article.URLs {
			if linkTag.ContentType == "application/pdf" {
				articleURL = linkTag.Link
			}
		}

		newArticle := pb.Result{
			Id:            article.Id,
			Author:        strings.Join(article.Authors[:], ";"),
			Title:         article.Title,
			PublishedDate: article.PublishedDate,
			Publisher:     "arXiv",
			Description:   article.Description,
			Url:           articleURL,
		}
		result = append(result, &newArticle)
	}

	return result
}
