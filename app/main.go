package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	_"strconv" 

	"github.com/gocolly/colly"
)

func main() {
	allName := make([]string, 0);

	collector := colly.NewCollector(
		colly.AllowedDomains("tvsty.my", "www.tvsty.my"), 
	)

	selector := "#post-5985 > div > div.wp-block-cover.alignfull.is-light.has-parallax.ticss-abaa7e82 > div > section > div > div > div > div > div > div.post-60898.movie.type-movie.status-publish.has-post-thumbnail.hentry.movie_genre-news-anchor.movie_genre-news-reporter > div.movie__body > div.movie__info > div.movie__info--head > a > h3";

	collector.OnHTML(selector, func(element *colly.HTMLElement) {
		// factId, err := strconv.Atoi(element.Attr("id"))
		// if err != nil {
		// 	log.Println("Could not get id")
		// }

		name := element.Text

		allName = append(allName, name);
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String()); 
	})

	collector.Visit("tvstv.my/our-host/");   

	fmt.Println(allName); 

	writeJSON(allName); 
}

func writeJSON(data []string) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file");
		return; 
	}

	_ = ioutil.WriteFile("namelist.json", file, 0644);
}