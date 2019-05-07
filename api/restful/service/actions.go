package service

import (
	"fmt"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/github"
	"github.com/MarkLux/GOLD/api/restful/orm"
	docker "github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"log"
	"os"
)

type Action struct {
	Type string
	FunctionService orm.FunctionService
	TargetBranch string
	TargetVersion string
	Operator orm.User
}

func (s FunctionService) buildImage(f orm.FunctionService, buildLog *orm.OperateLogs) (err error) {
	if f.GitHead == "" {
		gitCli := github.GithubClient{Maintainer: f.GitMaintainer, Repo: f.GitRepo}
		f.GitHead = gitCli.GetLastCommitSha(f.GitBranch)
	}
	err = s.updateStatus(f.Id, constant.ServiceStatusImageBuilding)
	if err != nil {
		log.Println("fail to update function service status, ", err)
		err = errors.GenUnknownError()
		return
	}
	// build args
	bArgs := make(map[string]*string)
	gitUrl := fmt.Sprintf("https://github.com/%s/%s", f.GitMaintainer, f.GitRepo)
	bArgs["REPO_URL"] = &gitUrl
	bArgs["REPO_NAME"] = &f.GitRepo
	bArgs["BRANCH"] = &f.GitBranch
	bArgs["COMMIT_ID"] = &f.GitHead
	// open build context
	bContext, err :=  os.Open(constant.DockerfilePath)
	if err != nil {
		_ = s.opService.FailOperateLog(buildLog, "fail to open docker context")
		log.Println("fail to open docker context, ", err)
		err = errors.GenUnknownError()
		return
	}
	// start build
	imgName := constant.DockerRegistry + "/" + f.ServiceName + ":" + f.GitHead
	rsp, err := s.dockerCli.ImageBuild(context.Background(), bContext, docker.ImageBuildOptions{
		Dockerfile: "tmp/Dockerfile",
		BuildArgs: bArgs,
		// ban cache for git
		NoCache: true,
		// name:tag
		Tags: []string{imgName},
	})
	if err != nil {
		_ = s.opService.FailOperateLog(buildLog, "fail to execute build command")
		log.Println("fail to image build, ", err)
		err = errors.GenUnknownError()
		return
	}
	defer rsp.Body.Close()
	err = s.opService.ContinueOperateLog(buildLog, ActionImgBuilding, rsp.Body, true)
	if err != nil {
		log.Println("fail to record output, ", err)
		// TODO check if the build succeed? there can be no api, maybe we have to read the output
	}
	// start pushing
	pRsp, err := s.dockerCli.ImagePush(context.Background(), imgName, docker.ImagePushOptions{
		RegistryAuth: "gold",
	})
	if err != nil {
		_ = s.opService.FailOperateLog(buildLog, "fail to push image")
		log.Println("fail to push image", err)
	}
	defer pRsp.Close()
	_ = s.opService.ContinueOperateLog(buildLog, ActionImgPushing, pRsp, true)
	return nil
}

func (s FunctionService) updateStatus(fId int64, status string) error {
	_, err := s.engine.Id(fId).Cols("status").Update(&orm.FunctionService{Status: status})
	return err
}