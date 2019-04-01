package constant

const (
	// the wrapper container would listen at default service port
	DefaultServicePort = "8080"
	// gold service namespace in k8s
	GoldNamespace = "gold"
	// default client timeout
	DefaultClientTimeOut = 3000
	// env var name of service name
	GoldServiceNameEnvKey = "GOLD_SERVICE_NAME"
	// redis service
	GoldRedisServiceName = "gold-service"
	// redis port
	GoldRedisServicePort = 6379
	// mongo primary entry
	GoldMongoPrimaryEndPoint = "mongod-0.mongo-service.gold.svc.cluster.local"
	// mongo port
	GoldMongoServicePort = 27017
)
