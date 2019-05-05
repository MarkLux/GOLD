package constant

const (
	DataBaseDriver = "mysql"
	DataBaseName = "gold"
	DataBaseUser = "root"
	DataBasePwd = "qwe123"

	RedisAddr = "localhost:6379"

	LoginTokenExpiredTime = 10800000

	KubeConfigPath = "/Users/lumin/.kube/config"

	DockerfilePath = "/Users/lumin/Projects/Go/GOLD/api/build/tmp.tar"

	DockerRegistry = "gold-registry:8099"
)

// service status
const (
	ServiceStatusCreated = "CREATED"
	ServiceStatusImageBuilding = "IMAGE_BUILDING"
	ServiceStatusImageBuildFail = "IMAGE_BUILD_FAIL"
	ServiceStatusPublishing = "PUBLISHING"
	ServiceStatusPublishFail = "PUBLISH_FAIL"
	ServiceStatusPublished = "PUBLISHED"
	ServiceStatusRollBacking = "ROLL_BACKING"
	ServiceStatusRollBackFail = "ROLL_BACK_FAIL"
	ServiceStatusRollBacked = "ROLL_BACKED"
	ServiceStatusDeleted = "DELETED"
)

// operate types
const (
	OperateBuild = "BUILD"
	OperatePublish = "PUBLISH"
	OperateRollBack = "ROLLBACK"
)