package bitbucket

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

type PullRequestPage struct {
	pagelen int				`json:"pagelen"`
	size int				`json:"size"`
	page int				`json:"page"`
	values []PullRequest	`json:"values"`
}

type Author struct {
	nickname string			`json:"nickname"`
	display_name string		`json:"display_name"`
}

type PullRequest struct {
	title string		`json:"title"`
	source Reference	`json:"source"`
	target Reference	`json:"target"`
	author Author		`json:"author"`
}

type Reference struct{
	branch string	`json:"branch"`
}

func ParsePr(body io.Reader) ([]PullRequest, error){

	res := map[string]interface{}{}
	err := utils.ParseJSON(&res, body)

	if err != nil {
		return nil, err
	}

	values := res["values"].([]interface{})
	parsed := make([]PullRequest, len(values))
	for i, value := range values {
		valueSeeker := value.(map[string]interface{})
		author :=valueSeeker["author"].(map[string]interface{})
		source :=valueSeeker["source"].(map[string]interface{})
		target :=valueSeeker["destination"].(map[string]interface{})
		parsed[i] = PullRequest{
			title: valueSeeker["title"].(string),
			author: Author{
				display_name: author["display_name"].(string),
				nickname: author["nickname"].(string),
			},
			source: Reference{
				branch:  source["branch"].(map[string]interface{})["name"].(string),
			},
			target: Reference{
				branch:  target["branch"].(map[string]interface{})["name"].(string),
			},
		}
	}

	return parsed, nil
}


func (widget *Widget) GetPrs() ([]PullRequest, error) {
	url := fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/pullrequests",
		widget.settings.workspace,
		widget.settings.repo,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", widget.settings.token))
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return ParsePr(resp.Body)
}
