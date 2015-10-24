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

const (
	ENV_API_KEY = "44642-ef80a8999e99444da2f6b65c"
)

func apiCredentials() string {
	return ENV_API_KEY // os.Getenv(ENV_API_KEY)
}

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
func GetPocketRequestToken(callbackUrl string) string {
	log.Println("creds")
	apiKey := apiCredentials()
	log.Println(apiKey)
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/request",
		url.Values{"consumer_key": {apiKey}, "redirect_uri": {callbackUrl}},
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

func GetPocketAccessToken(code string, callbackUrl string) (string, string) {
	apiKey := apiCredentials()

	browser.OpenURL("https://getpocket.com/auth/authorize?request_token=" + code + "&redirect_uri=" + callbackUrl)
	time.Sleep(time.Millisecond * 10000)
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/authorize",
		url.Values{"consumer_key": {apiKey}, "code": {code}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)
	log.Println(err)
	log.Println(values)
	return values.Get("username"), values.Get("access_token")

}

func AddItemToPocket(access_token string, screenname string, tweet_id int64) {
	apiKey := apiCredentials()
	resp, err := http.PostForm(
		"https://getpocket.com/v3/add",
		url.Values{"consumer_key": {apiKey}, "access_token": {access_token}, "url": {"https://twitter.com/" + screenname + "/status/" + strconv.FormatInt(tweet_id, 10)}, "tweet_id": {strconv.FormatInt(tweet_id, 10)}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)
	log.Println(err)
	log.Println(values)
}
