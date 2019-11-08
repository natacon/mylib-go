package redmine

import (
	"encoding/json"
	"fmt"
	"github.com/natacon/mylib-go/common"
	"log"
	"os"
	"regexp"
)

var (
	RedmineUrl string = "http://192.168.27.200/redmine"
	key        string = os.Getenv("REDMINE_APIKEY")
	re                = regexp.MustCompile(`\d{4,}`)
)

type Redmine struct {
	Endpoint string
	ApiKey   string
}

func (r *Redmine) CreateIssueUrls(issueIds []string) (urls []string) {
	for _, issueId := range issueIds {
		url := fmt.Sprintf("%s/issues/%s.json?key=%s", r.Endpoint, issueId, r.ApiKey)
		if re.MatchString(issueId) {
			log.Printf("url: %s\n", url)
			urls = append(urls, url)
		}
	}
	return
}

func (r *Redmine) SelectIssues(urls []string) (issues []Issue, err error) {
	for _, url := range urls {
		body, err := common.CallApi("GET", url)
		if err != nil {
			return issues, err
		}
		var issue Issue
		issueBody := body[9 : len(body)-1] // issue本体部分にjson文字列を限定する TODO:やめる
		if err := json.Unmarshal(issueBody, &issue); err != nil {
			return issues, err
		} else {
			issues = append(issues, issue)
		}
	}
	return
}

func CreateMessage(issues []Issue) (message string) {
	for _, issue := range issues {
		message += fmt.Sprintf("%s %s #%d: %s\n%s/issues/%d\n\n",
			issue.Status.Name, issue.Tracker.Name, issue.ID, issue.Subject, RedmineUrl, issue.ID)
	}
	return
}

func CreateUrls(m []string) (urls []string) {
	for _, e := range m {
		url := fmt.Sprintf("%s/issues/%s.json?key=%s", RedmineUrl, e, key)
		if re.MatchString(e) {
			log.Printf("url: %s\n", url)
			urls = append(urls, url)
		}
	}
	return
}

func GetIssues(urls []string) (issues []Issue, err error) {
	for _, url := range urls {
		body, err := common.CallApi("GET", url)
		if err != nil {
			return issues, err
		}
		var issue Issue
		issueBody := body[9 : len(body)-1] // issue本体部分にjson文字列を限定する TODO:やめる
		if err := json.Unmarshal(issueBody, &issue); err != nil {
			return issues, err
		} else {
			issues = append(issues, issue)
		}
	}
	return
}

// TODO: 未実装
func extractIssueId(text string) (ids []string) {
	// TODO: メッセージから数値4桁を全て取得する（それがいいかは別として）
	//texts = re.findall(r'(?<![0-9])[0-9]{4}(?![0-9])', message.body['text'])
	re := regexp.MustCompile(`(?<![0-9])[0-9]{4}(?![0-9])`)
	all := re.FindAllString(text, -1)

	// TODO: 重複削除
	//texts = list(set(texts))
	m := make(map[string]struct{})
	ids = make([]string, 0)
	for _, e := range all {
		if _, ok := m[e]; !ok {
			m[e] = struct{}{}
			ids = append(ids, e)
		}
	}
	//texts.sort()
	return
}
