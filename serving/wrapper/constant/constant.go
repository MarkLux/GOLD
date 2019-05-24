package constant

const (
	// the wrapper container would listen at default restful port
	DefaultServicePort = "8080"
	// gold restful namespace in k8s
	GoldNamespace = "gold"
	// default client timeout
	DefaultClientTimeOut = 3000
	// env var name of restful name
	GoldServiceNameEnvKey = "GOLD_SERVICE_NAME"
	// redis restful
	GoldRedisServiceName = "gold-restful"
	// redis port
	GoldRedisServicePort = "6379"
	// mongo primary entry
	GoldMongoPrimaryEndPoint = "mongod-0.mongo-service.gold.svc.cluster.local"
	// mongo port
	GoldMongoServicePort = "27017"
)
