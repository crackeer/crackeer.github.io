package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func NewIndex() (bleve.Index, error) {
	index, err := bleve.Open("example")
	if err != nil {
		mapping := bleve.NewIndexMapping()
		blog := bleve.NewDocumentMapping()
		nameFieldMapping := bleve.NewTextFieldMapping()
		nameFieldMapping.Analyzer = "en"
		nameFieldMapping.IncludeInAll = true
		nameFieldMapping.Name = "s"

		blog.AddFieldMappingsAt("name", nameFieldMapping)
		mapping.AddDocumentMapping("beer", blog)
		return bleve.New("example", mapping)
	}
	return index, nil
}

func main() {
	aboutSimple()
}

func aboutSimple() {
	// open a new index

	index, err := NewIndex()
	if err != nil {
		panic(err)
	}

	/*
		data := map[string]interface{}{
			"abc":  78,
			"name": "simple",
		}

		// index some data
		index.Index("id11", data)
		fields, err := index.Fields()
		fmt.Println(fields, err)
	*/

	// search for some text
	query := bleve.NewMatchQuery("simple")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range searchResults.Hits {
		fmt.Println(item.ID)
	}

	fmt.Println(searchResults.Hits)
}
