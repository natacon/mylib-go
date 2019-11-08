package redmine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var (
	RedmineURL string = "http://192.168.27.200/redmine"
	key        string = os.Getenv("REDMINE_APIKEY")
	re                = regexp.MustCompile(`\d{4,}`)
)

func CreateMessage(issues []Issue) (message string) {
	for _, issue := range issues {
		message += fmt.Sprintf("%s %s #%d: %s\n%s/issues/%d\n\n",
			issue.Status.Name, issue.Tracker.Name, issue.ID, issue.Subject, RedmineURL, issue.ID)
	}
	return
}

func GetURLs(m []string) (urls []string) {
	for _, e := range m {
		url := fmt.Sprintf("%s/issues/%s.json?key=%s", RedmineURL, e, key)
		if re.MatchString(e) {
			log.Printf("url: %s\n", url)
			urls = append(urls, url)
		}
	}
	return
}

func GetIssues(urls []string) (issues []Issue) {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Printf("[ERROR] Faild to request: %v", err)
		} else if res.StatusCode >= 400 {
			log.Printf("[ERROR] Faild to request: %d", res.StatusCode)
		} else {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("[ERROR] Faild to read response body: %v", err)
			} else {
				var issue Issue
				issueBody := body[9 : len(body)-1] // issue本体部分にjson文字列を限定する
				if err := json.Unmarshal(issueBody, &issue); err != nil {
					log.Printf("[ERROR] Faild to unmarshal: %v", err)
				} else {
					issues = append(issues, issue)
				}
			}
		}
		_ = res.Body.Close()
	}
	return
}

// TODO: 未実装
func extractIssueId(text string) (ids []string) {
	// TODO: メッセージから数値4桁を全て取得する
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
