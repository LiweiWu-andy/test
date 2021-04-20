package main

import (
	"fmt"
	"github.com/ParticleMedia/social-comment/parser/fetcher"
	sj "github.com/bitly/go-simplejson"
	"io/ioutil"
	"strings"
)

type Comment struct {
	ID				string		`json:"id"`
	Message			string		`json:"message"`
	LikeCount		int			`json:"like_count"`
	CreateTS		string		`json:"create_ts"`
	CommentCount	int			`json:"comment_count"`
}
type Fetcher = fetcher.Fetcher

var (
	gCommentFields  = []string{"like_count", "comment_count", "created_time", "message", "attachment", "from", "comments"}
	gPostsUrl 	 = "https://graph.facebook.com/v3.2/%s/comments?fields=%s&order=%s&limit=%d&access_token=%s"
	gAccessToken = "153277661752118|mpAYNH_jzO-jwO_PhGexz-xljgg"

)
const (
	commentOrder = "reverse_chronological"
)

type Parser interface {
	Parse(url string, limit int) (error)
	Platform() string
}

type FacebookParser struct {

}
func Fetch(url string) ([]byte, error) {
	gFetcher := fetcher.NewDefaultFetcher()
	fmt.Println(url)
	if resp, err := gFetcher.Get(url); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}

func fetchCommentByPostID(PostID string, num int) ([]byte, error) {
	if strings.Contains(PostID, "-") {
		ss := strings.Split(PostID, "-")
		PostID = ss[len(ss)-1]
	}

	fields := strings.Join(gCommentFields, ",")
	pageUrl := fmt.Sprintf(gPostsUrl, PostID, fields, commentOrder, num, gAccessToken)
	html, err := Fetch(pageUrl)
	if err != nil {
		return nil, fmt.Errorf("fetch pageurl error %s", err)
	}
	return html, nil
}

func ParseComment (html []byte) ([]Comment, error){
	comments := make([]Comment, 0)
	jsonObj, err := sj.NewJson(html)
	if err != nil {
		return nil, fmt.Errorf("new json error %s", err)
	}
	data := jsonObj.Get("data")
	for i := range data.MustArray() {
		comments = append(comments, Comment{
			data.GetIndex(i).Get("id").MustString(),
			data.GetIndex(i).Get("message").MustString(),
			data.GetIndex(i).Get("like_count").MustInt(),
			data.GetIndex(i).Get("created_time").MustString(),
			data.GetIndex(i).Get("comment_count").MustInt(),
		})
	}
	return comments, nil
}

func getCommentByPostID(PostID string, num int) ([]Comment, error) {
	html, err := fetchCommentByPostID("10156401171628788", 5)
	if err != nil {
		return nil, err
	}
	comments, err := ParseComment(html)
	if err != nil {
		return nil, err
	}
	return comments, nil
}