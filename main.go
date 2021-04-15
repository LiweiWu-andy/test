package main

import (
	"fmt"
	"github.com/ParticleMedia/social-fetch/parser/fetcher"
	"io/ioutil"
	"strings"
)

type Fetcher = fetcher.Fetcher

var (
	gFetcher Fetcher
	gPageFields  = []string{"like_count", "comment_count", "created_time", "message", "attachment", "from", "comments"}
	// gPostsUrl2    = "https://graph.facebook.com/v3.2/%s?fields=%s&access_token=%s"
	gPostsUrl 	 = "https://graph.facebook.com/v3.2/%s/comments?fields=%s&order=%s&limit=%d&access_token=%s"
	gAccessToken = "153277661752118|mpAYNH_jzO-jwO_PhGexz-xljgg"

	)
func Fetch(url string) ([]byte, error) {
	fmt.Println(url)
	if resp, err := gFetcher.Get(url); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}

func getAllFieldsPage(userName string, num int) ([]byte, error) {
	if strings.Contains(userName, "-") {
		ss := strings.Split(userName, "-")
		userName = ss[len(ss)-1]
	}

	fields := strings.Join(gPageFields, ",")
	pageUrl := fmt.Sprintf(gPostsUrl, userName, fields, "reverse_chronological", 3, gAccessToken)
	html, err := Fetch(pageUrl)
	if err != nil {
		return nil, fmt.Errorf("fetch pageurl error %s", err)
	}
	return html, nil
}

func main() {
	gFetcher = fetcher.NewDefaultFetcher()
	html, _ := getAllFieldsPage("10156401171628788", 3)
	fmt.Println(string(html))
}
