package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/chimeracoder/anaconda"
	"github.com/jchannon/PoGo/pocket"
	"github.com/jchannon/PoGo/twitter"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print("go run main.go")
	fmt.Print("  --consumerkey <consumerkey>")
	fmt.Print("  --consumersecret <consumersecret>")
	fmt.Println("  --pocketapikey <pocketapikey>")
	fmt.Println("")
	fmt.Println("In order to get your consumerkey and consumersecret, you must register an 'app' at twitter.com:")
	fmt.Println("https://dev.twitter.com/apps/new")
	fmt.Println("In order to get your pocketapikey, you must register an 'app' at getpocket.com:")
	fmt.Println("https://getpocket.com/developer/")
}

func startWebServer() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}

func main() {
	var consumerKey = flag.String(
		"consumerkey",
		"",
		"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")

	var consumerSecret = flag.String(
		"consumersecret",
		"",
		"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")

	var apiKey = flag.String(
		"pocketapikey",
		"",
		"API key from Pocket")

	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 || len(*apiKey) == 0 {
		fmt.Println("You must set the --consumerkey and --consumersecret and --pocketapi flags.")
		fmt.Println("---")
		usage()
		os.Exit(1)
	}

	twitter := &twitter.Twitter{}
	favourites, err := twitter.GetFavourites(consumerKey, consumerSecret)
	if err != nil {
		log.Fatal(err)
	}

	favouriteLength := len(favourites)

	for favouriteLength == 200 {
		maxid := strconv.FormatInt(favourites[len(favourites)-1].Id, 10)
		pagedFavourites, err := twitter.GetPagedFavourites(consumerKey, consumerSecret, twitter.Token, twitter.Secret, maxid)
		if err != nil {
			log.Fatal(err)
		}
		favourites = append(favourites, pagedFavourites...)

		favouriteLength = len(favourites)
	}

	go startWebServer()

	data := pocket.GetPocketRequestToken(apiKey, "http://localhost:3000/")
	pocket.AuthorizePocket(data, "http://localhost:3000/")
	// fmt.Println(data)
	fmt.Println("(4) Press Enter when authorized with Pocket.")
	instr := ""
	fmt.Scanln(&instr)
	_, pocketaccesstoken := pocket.GetPocketAccessToken(apiKey, data, "http://yahoo.co.uk")

	for _, tweet := range favourites {

		if len(tweet.Entities.Urls) > 0 {
			for _, tweeturl := range tweet.Entities.Urls {

				url, err := url.Parse(tweeturl.Expanded_url)
				if err != nil {

				}
				ext := filepath.Ext(url.Path)
				if len(ext) == 0 || ext == ".html" {
					addUrlInTweetToPocket(apiKey, pocketaccesstoken, tweeturl.Expanded_url, tweet.Id)
				} else {
					//addBasicTweetToPocket(apiKey, pocketaccesstoken, tweet)
					fmt.Println("Not processed : " + tweeturl.Expanded_url)
				}
			}
		} else {
			addBasicTweetToPocket(apiKey, pocketaccesstoken, tweet)
		}

		break

	}

}

func addBasicTweetToPocket(apiKey *string, accesstoken string, tweet anaconda.Tweet) {
	pocketurl := "https://twitter.com/" + tweet.User.ScreenName + "/status/" + strconv.FormatInt(tweet.Id, 10)
	fmt.Println("Adding to pocket : " + pocketurl)
	pocket.AddItemToPocket(apiKey, accesstoken, pocketurl, tweet.Id)
}

func addUrlInTweetToPocket(apiKey *string, accesstoken string, urlintweet string, Id int64) {
	fmt.Println("Adding to pocket : " + urlintweet)
	pocket.AddItemToPocket(apiKey, accesstoken, urlintweet, Id)
}
