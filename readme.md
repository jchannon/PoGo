#gofavpocket

Written in Golang, this tool imports Twitter favourites into Pocket.

###Instructions

* Get yourself a Twitter access token and secret from http://dev.twitter.com
* Get yourself a Pocket access token from https://getpocket.com/developer/
* Execute the app `go run main.go --consumerkey [TWITTER KEY] --consumersecret [TWITTER SECRET] --pocketapikey [POCKETAPIKEY]`

**Notes:**

This will read tweets marked as favourites.  If the tweet contains no links in the tweet, it will add the link to the tweet into Pocket.  If the tweet contains a link it will add that to Pocket **only** if the URL has no file extension **OR** if the URL has a `.html` extension.  If the tweet contains a link to something other than the above condition it will be logged to the console that it has not been added to Pocket.
![](https://cloud.githubusercontent.com/assets/105126/10717526/041bfb34-7b53-11e5-92df-4d82934a1645.png)
