/*

Takes care of OAuth to the Pocket service.
To use, need to set Pocket API consumer key as environment variable POCKET_API_KEY.

*/
package pocket

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/browser"
)

func responseBodyAsValues(r *http.Response) (url.Values, error) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return url.Values{}, err
	}

	return url.ParseQuery(string(body))
}

// GetPocketRequestToken will get the first request token, kicks off the authentication
// process.
func GetPocketRequestToken(apiKey *string, callbackUrl string) string {
	apikeyValue := *apiKey
	log.Println("creds")
	log.Println(apiKey)
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/request",
		url.Values{"consumer_key": {apikeyValue}, "redirect_uri": {callbackUrl}},
	)
	log.Println("got creds")
	log.Println(err)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)
	log.Println(err)
	log.Println(values)
	log.Println(values.Get("code"))

	return values.Get("code")
}

func GetPocketAccessToken(apiKey *string, code string, callbackUrl string) (string, string) {
	apikeyValue := *apiKey
	browser.OpenURL("https://getpocket.com/auth/authorize?request_token=" + code + "&redirect_uri=" + callbackUrl)
	time.Sleep(time.Millisecond * 7000)
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/authorize",
		url.Values{"consumer_key": {apikeyValue}, "code": {code}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)
	log.Println(err)
	log.Println(values)
	return values.Get("username"), values.Get("access_token")

}

func AddItemToPocket(apiKey *string, access_token string, tweeturl string, tweet_id int64) {
	apikeyValue := *apiKey
	resp, err := http.PostForm(
		"https://getpocket.com/v3/add",
		url.Values{"consumer_key": {apikeyValue}, "access_token": {access_token}, "url": {tweeturl}, "tweet_id": {strconv.FormatInt(tweet_id, 10)}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)
	log.Println(err)
	log.Println(values)
}
