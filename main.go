package main

import (
	"fmt"
	"encoding/json"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/cyarie/trump_tweets/config"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	dev := viper.Sub("app.development.twitter")
	fmt.Println(dev)
	fmt.Println(dev.Get("consumerKey"))

	consumerKey := viper.Get("app.development.twitter.consumerKey").(string)
	consumerSecret := viper.Get("app.development.twitter.consumerSecret").(string)
	tokenKey := viper.Get("app.development.twitter.token").(string)
	tokenSecret := viper.Get("app.development.twitter.tokenSecret").(string)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(tokenKey, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	userShowParams := &twitter.UserShowParams{ScreenName: "realDonaldTrump"}
	user, _, _ := client.Users.Show(userShowParams)
	fmt.Printf("USERS SHOW:\n%v\n", user)

	f := new(bool)
	*f = false
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: "realDonaldTrump", Count: 1, IncludeRetweets: f}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
	for tweet := range tweets {
		t, _ := json.MarshalIndent(tweets[tweet], "", "    ")
		fmt.Println(string(t))
	}
}


