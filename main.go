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
	gPageFields  = []string{"category", "category_list", "about", "website", "name", "id", "overall_star_rating",
		"posts.limit(%d){properties,likes.limit(0).summary(true),comments.limit(0).summary(true),message,description,created_time,type,object_id,link,attachments,shares,status_type,parent_id}",
		"birthday", "artists_we_like", "bio", "band_interests", "is_community_page", "release_date", "is_verified", "talking_about_count",
		"starring", "season", "produced_by", "genre", "directed_by", "network", "description", "engagement", "fan_count", "link", "location",
		"username", "verification_status", "picture.type(large)", "emails", "checkins", "price_range", "hours", "phone", "rating_count"}
	gPostsUrl    = "https://graph.facebook.com/v3.2/%s?fields=%s&access_token=%s"
	gAccessToken = "153277661752118|mpAYNH_jzO-jwO_PhGexz-xljgg"

	)
func Fetch(url string) ([]byte, error) {
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

	fields := fmt.Sprintf(strings.Join(gPageFields, ","), num)
	pageUrl := fmt.Sprintf(gPostsUrl, userName, fields, gAccessToken)
	html, err := Fetch(pageUrl)
	if err != nil {
		return nil, fmt.Errorf("fetch pageurl error %s", err)
	}
	return html, nil
}

func main() {
	gFetcher = fetcher.NewDefaultFetcher()
	html, _ := getAllFieldsPage("https://www.facebook.com/1172972172724188", 3)
	fmt.Println(string(html))
}
