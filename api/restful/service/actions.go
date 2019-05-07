package service

import (
	"fmt"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/github"
	"github.com/MarkLux/GOLD/api/restful/orm"
	docker "github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	appV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	asV1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
)

type Action struct {
	Type            string
	FunctionService orm.FunctionService
	TargetBranch    string
	TargetVersion   string
	Operator        orm.User
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
	bContext, err := os.Open(constant.DockerfilePath)
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
		BuildArgs:  bArgs,
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

	return
}

func (s FunctionService) pushImage(f orm.FunctionService, buildLog *orm.OperateLogs) (err error) {
	imgName := constant.DockerRegistry + "/" + f.ServiceName + ":" + f.GitHead
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
	return
}

func (s FunctionService) initK8sService(f orm.FunctionService, opLog *orm.OperateLogs) (err error) {
	_ = s.opService.ContinueOperateLog(opLog, ActionPublishing, nil, false)
	// prepare
	img := constant.GoldRegistry + "/" + f.ServiceName + ":" + f.GitHead
	labelMap := map[string]string{"app": f.ServiceName}

	// 1. create k8s deployment
	dp := &appV1.Deployment{}
	dp.Name = f.ServiceName
	dp.Namespace = constant.GoldNameSpace
	container := coreV1.Container{
		Name:  f.ServiceName,
		Image: img,
		Ports: []coreV1.ContainerPort{{Name: "rpcPort", ContainerPort: constant.RpcPort}},
		Resources: coreV1.ResourceRequirements{
			Limits:   coreV1.ResourceList{"cpu": resource.MustParse(constant.LimitCpu), "memory": resource.MustParse(constant.LimitMem)},
			Requests: coreV1.ResourceList{"cpu": resource.MustParse(constant.RequestCpu), "memory": resource.MustParse(constant.RequestMem)},
		},
	}
	// default replicas num: min instance.
	replicas := int32(f.MinInstance)
	dp.Spec.Replicas = &replicas
	dp.Spec.Template.Labels = labelMap
	dp.Spec.Template.Spec.Containers = []coreV1.Container{container}
	_, err = s.k8sCli.AppsV1().Deployments(constant.GoldNameSpace).Create(dp)
	if err != nil {
		log.Println("fail to create deployment, ", err)
		_ = s.opService.FailOperateLog(opLog, "fail to create deployment")
		return
	}
	// 2. create k8s service
	svc := &coreV1.Service{}
	svc.Name = f.ServiceName
	svc.Namespace = constant.GoldNameSpace
	svc.Labels = labelMap
	svc.Spec.Selector = labelMap
	svc.Spec.Type = coreV1.ServiceTypeNodePort
	svc.Spec.Ports = []coreV1.ServicePort{{
		Name:       "rpcPort",
		Protocol:   coreV1.ProtocolTCP,
		Port:       constant.RpcPort,
		TargetPort: constant.RpcPort,
	}}
	_, err = s.k8sCli.CoreV1().Services(constant.GoldNameSpace).Create(svc)
	if err != nil {
		log.Println("fail to create service, ", err)
		_ = s.opService.FailOperateLog(opLog, "fail to create service")
	}
	// 3. config HPA
	hpa := &asV1.HorizontalPodAutoscaler{}
	hpa.Name = f.ServiceName
	hpa.Namespace = constant.GoldNameSpace
	hpa.Labels = labelMap
	hpa.Spec.ScaleTargetRef = asV1.CrossVersionObjectReference{
		Kind: "Deployment",
		Name: f.ServiceName,
		APIVersion: "extensions/v1beta1",
	}
	minIns := int32(f.MinInstance)
	maxIns := int32(f.MaxInstance)
	hpa.Spec.MinReplicas = &minIns
	hpa.Spec.MaxReplicas = maxIns
	_, err = s.k8sCli.AutoscalingV1().HorizontalPodAutoscalers(constant.GoldNameSpace).Create(hpa)
	if err != nil {
		log.Println("fail to attach hpa, ", err)
		_ = s.opService.FailOperateLog(opLog, "fail to create hpa")
	}

	return
}

func (s FunctionService) publishK8sService(f orm.FunctionService, opLog *orm.OperateLogs) (err error) {
	_ = s.opService.ContinueOperateLog(opLog, ActionPublishing, nil, false)
	img := constant.GoldRegistry + "/" + f.ServiceName + ":" + f.GitHead
	// change the dp image, fire!
	dp, err := s.k8sCli.AppsV1().Deployments(constant.GoldNameSpace).Get(f.ServiceName, v1.GetOptions{})
	if err != nil {
		log.Println("fail to get deployment resource, ", err)
		_ = s.opService.FailOperateLog(opLog, "fail to locate deployment")
		return

	}
	dp.Spec.Template.Spec.Containers[0].Image = img
	_, err = s.k8sCli.AppsV1().Deployments(constant.GoldNameSpace).Update(dp)
	if err != nil {
		log.Println("fail to update deployment image, ", err)
		_ = s.opService.FailOperateLog(opLog, "fail to update deployment image")
		return
	}
	return
}

func (s FunctionService) updateStatus(fId int64, status string) error {
	_, err := s.engine.Id(fId).Cols("status").Update(&orm.FunctionService{Status: status})
	return err
}
