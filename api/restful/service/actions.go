package service

import (
	"fmt"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/github"
	"github.com/MarkLux/GOLD/api/restful/orm"
	docker "github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func (s FunctionService) buildImage(f orm.FunctionService, buildLog orm.OperateLogs) (err error) {
	if f.GitHead == "" {
		// default use last commit as the version
		gitCli := github.GithubClient{Maintainer: f.GitMaintainer, Repo: f.GitRepo}
		f.GitHead = gitCli.GetLastCommitSha(f.GitBranch)
	}
	err = s.updateStatus(f.Id, constant.ServiceStatusImageBuilding)
	if err != nil {
		log.Println("fail to update function service status, ", err)
		err = errors.GenUnknownError()
		return
	}
	s.updateOperate(buildLog.Id, "BUILDING", buildLog.Log, false)
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
		s.failOperate(buildLog.Id, "fail to open docker context")
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
		s.failOperate(buildLog.Id, "fail to execute build command")
		log.Println("fail to image build, ", err)
		err = errors.GenUnknownError()
		return
	}
	defer rsp.Body.Close()
	// record the output
	outputBytes, _ := ioutil.ReadAll(rsp.Body)
	output := string(outputBytes)
	output += "\n build succeed, start pushing...\n"
	// build succeed, update log
	s.updateOperate(buildLog.Id, "PUSHING", output, false)
	// start pushing
	pRsp, err := s.dockerCli.ImagePush(context.Background(), imgName, docker.ImagePushOptions{
		RegistryAuth: "gold",
	})
	if err != nil {
		output += fmt.Sprintf("\n fail to push image, get error: %s", err)
		s.failOperate(buildLog.Id, output)
		log.Println("fail to push image", err)
	}
	defer pRsp.Close()
	pushOutputBytes, _ := ioutil.ReadAll(pRsp)
	output += string(pushOutputBytes)
	s.updateOperate(buildLog.Id, "PUSHED", output, false)

	return nil
}

func (s FunctionService) updateStatus(fId int64, status string) error {
	_, err := s.engine.Id(fId).Cols("status").Update(&orm.FunctionService{Status: status})
	return err
}

func (s FunctionService) failOperate(opId int64, log string) {
	current := time.Now().Unix()
	s.engine.Table(orm.OperateLogs{}).
		ID(opId).Update(&orm.OperateLogs{
		CurrentAction: "FAILED",
		Log: log,
		End: current,
		Update: current,
	})
}

func (s FunctionService) updateOperate(opId int64, act string, log string, end bool) {
	current := time.Now().Unix()
	update := &orm.OperateLogs{
		CurrentAction: act,
		Update: current,
		Log: log,
	}
	if end {
		update.End = current
	}
	s.engine.Table(orm.OperateLogs{}).ID(opId).Update(update)
}