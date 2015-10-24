package main

import (
	//"flag"
	"flag"
	"fmt"
	"log"
	"os"
	//"log"
	//"os"

	"github.com/jchannon/gofavpocket/pocket"
	"github.com/jchannon/gofavpocket/twitter"
	//"github.com/jchannon/gofavpocket/twitter"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Print("go run examples/twitter/twitter.go")
	fmt.Print("  --consumerkey <consumerkey>")
	fmt.Println("  --consumersecret <consumersecret>")
	fmt.Println("")
	fmt.Println("In order to get your consumerkey and consumersecret, you must register an 'app' at twitter.com:")
	fmt.Println("https://dev.twitter.com/apps/new")
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

	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 {
		fmt.Println("You must set the --consumerkey and --consumersecret flags.")
		fmt.Println("---")
		usage()
		os.Exit(1)
	}

	blah := &twitter.Twitter{}
	favourites, err := blah.GetFavourites(consumerKey, consumerSecret)
	if err != nil {
		log.Fatal(err)
	}

	data := pocket.GetPocketRequestToken("http://google.co.uk")

	fmt.Println(data)
	_, kkk := pocket.GetPocketAccessToken(data, "http://yahoo.co.uk")

	for _, tweet := range favourites {

		pocket.AddItemToPocket(kkk, tweet.User.ScreenName, tweet.Id)
		break
	}

}
