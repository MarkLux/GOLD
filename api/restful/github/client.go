package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubClient struct {
	Maintainer string
	Repo string
}

func (c GithubClient) GetLastCommitSha(branch string) string {
	apiUrlPattern := "https://api.github.com/repos/%s/%s/commits/%s"
	apiUrl := fmt.Sprintf(apiUrlPattern, c.Maintainer, c.Repo, branch)
	restRsp, err := http.Get(apiUrl)
	if err != nil {
		return ""
	}
	defer restRsp.Body.Close()
	body, err := ioutil.ReadAll(restRsp.Body)
	if err != nil {
		return ""
	}
	v := make(map[string]interface{})
	err = json.Unmarshal(body, &v)
	if err != nil {
		return ""
	}
	return v["sha"].(string)
}