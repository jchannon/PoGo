package twitter

import (
	"fmt"
	"log"
	"net/url"

	"github.com/chimeracoder/anaconda"
	"github.com/mrjones/oauth"
)

//Twitter type
type Twitter struct {
	Token  string
	Secret string
}

//GetFavourites method
func (t *Twitter) GetFavourites(consumerKey *string, consumerSecret *string) ([]anaconda.Tweet, error) {

	c := oauth.NewConsumer(
		*consumerKey,
		*consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	c.Debug(false) //print out request contents

	requestToken, u, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + u)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	t.Token = accessToken.Token
	t.Secret = accessToken.Secret

	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(accessToken.Token, accessToken.Secret)
	v := url.Values{}
	v.Set("count", "200") //200 is max twitter accepts
	return api.GetFavorites(v)

}

func (t *Twitter) GetPagedFavourites(consumerKey *string, consumerSecret *string, apiToken string, apiSecret string, maxid string) ([]anaconda.Tweet, error) {
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(apiToken, apiSecret)

	v := url.Values{}
	v.Set("count", "200")
	v.Set("max_id", maxid)
	return api.GetFavorites(v)
}
