package constant

const (
	DataBaseDriver = "mysql"
	DataBaseName = "gold"
	DataBaseUser = "root"
	DataBasePwd = "root"

	RedisAddr = "localhost:6379"

	LoginTokenExpiredTime = 10800000

	RpcPort = 8080

	KubeConfigPath = "/Users/pbase1/.kube/config"
	DockerfilePath = "/Users/pbase1/Projects/Go/GOLD/api/build/tmp.tar"
	DockerRegistry = "gold-registry:5000"
	GoldRegistry = "gold-registry:5000"
	GoldNameSpace = "gold"
)

// service status
const (
	ServiceStatusCreated = "CREATED"
	ServiceStatusImageBuilding = "IMAGE_BUILDING"
	ServiceStatusImageBuildFail = "IMAGE_BUILD_FAIL"
	ServiceStatusImagePushing = "IMAGE_PUSHING"
	ServiceStatusImagePushFail = "IMAGE_PUSH_FAIL"
	ServiceStatusPublishing = "PUBLISHING"
	ServiceStatusPublishFail = "PUBLISH_FAIL"
	ServiceStatusPublished = "PUBLISHED"
	ServiceStatusRollBacking = "ROLL_BACKING"
	ServiceStatusRollBackFail = "ROLL_BACK_FAIL"
	ServiceStatusRollBacked = "ROLL_BACKED"
	ServiceStatusDeleted = "DELETED"
)

// hpa limits
const (
	LimitCpu = "25m"
	LimitMem = "128Mi"
	RequestCpu = "5m"
	RequestMem = "64Mi"
)

// operate types
const (
	OperateBuild = "BUILD"
	OperatePublish = "PUBLISH"
	OperateRollBack = "ROLLBACK"
)

// roles
const (
	RoleDev = "DEV"
	RoleAdmin = "ADMIN"
)