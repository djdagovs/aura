// +build pvfm

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/PonyvilleFM/aura/pvfm/pvl"
	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func genFname(u *discordgo.User) (string, error) {
	return fmt.Sprintf("%s - %s.mp3", u.Username, time.Now().Format(time.RFC822)), nil
}

var (
	consumerKey    = os.Getenv("TWITTER_KEY1")
	consumerSecret = os.Getenv("TWITTER_KEY2")

	accessToken  = os.Getenv("TWITTER_KEY3")
	accessSecret = os.Getenv("TWITTER_KEY4")
)

func newTwitter() *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	return twitter.NewClient(config.Client(oauth1.NoContext, token))
}

func announce(msg string) error {
	t := newTwitter()
	_, _, err := t.Statuses.Update(msg, nil)
	if err != nil {
		return err
	}

	return nil
}

func showName() (string, error) {
	cal, err := pvl.Get()
	if err != nil {
		return "", nil
	}

	now := cal.Result[0]
	return now.Title, nil
}
