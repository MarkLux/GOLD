package main

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/client"
	docker "github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	maintainer = "MarkLux"
	privateRegistry = "gold-registry:5000"
)

func main() {

	repoUrl := "https://github.com/MarkLux/GOLD-Bootstrap"
	repoName := "GOLD-Bootstrap"
	branchName := "demo"

	// create docker client

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	f , err:= os.Open("/Users/lumin/Projects/Go/GOLD/api/build/tmp.tar")
	if err != nil {
		panic(err)
	}

	// github api
	sha, e := getLatestCommitSha(repoName, maintainer, "demo")
	if e != nil {
		panic(e)
	}

	bArgs := make(map[string]*string)
	bArgs["REPO_URL"] = &repoUrl
	bArgs["REPO_NAME"] = &repoName
	bArgs["BRANCH"] = &branchName
	bArgs["COMMIT_ID"] = &sha

	imgName := privateRegistry +"/hello-restful:" + sha

	rsp, err :=cli.ImageBuild(ctx, f, docker.ImageBuildOptions{
		// this param is only suggest the file name of dockerfile, not path
		Dockerfile: "tmp/Dockerfile",
		BuildArgs: bArgs,
		// ban cache for git
		NoCache: true,
		// name:tag
		Tags: []string{imgName},
	})

	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()
	_, err = io.Copy(os.Stdout, rsp.Body)

	fmt.Println("build completed, start push.")

	pRsp, err := cli.ImagePush(ctx, imgName, docker.ImagePushOptions{
		RegistryAuth: "gold",
	})
	if err != nil {
		panic(err)
	}
	defer pRsp.Close()
	_, err = io.Copy(os.Stdout, pRsp)
}

func getLatestCommitSha(repoName string, maintainer string, branch string) (sha string, err error) {
	apiUrlPattern := "https://api.github.com/repos/%s/%s/commits/%s"
	apiUrl := fmt.Sprintf(apiUrlPattern, maintainer, repoName, branch)
	restRsp, err := http.Get(apiUrl)
	if err != nil {
		return
	}
	defer restRsp.Body.Close()
	body, err := ioutil.ReadAll(restRsp.Body)
	if err != nil {
		return
	}
	v := make(map[string]interface{})
	err = json.Unmarshal(body, &v)
	if err != nil {
		return
	}
	return v["sha"].(string), nil
}
