package env

var AppEnv string

const (
	ProdEnv = "prod"
	TestEnv = "test"
	DevEnv  = "dev"
)

func IsProd() bool {
	return AppEnv == "prod" || AppEnv == "prod_docker"
}

func IsTest() bool {
	return AppEnv == "test" || AppEnv == "test_docker"
}

func IsDev() bool {
	return !IsProd() && !IsTest()
}

func IsDocker() bool {
	return AppEnv == "prod_docker" || AppEnv == "test_docker" || AppEnv == "dev_docker"
}

func GetEnv() string {
	if IsProd() {
		return ProdEnv
	}
	if IsTest() {
		return TestEnv
	}
	return DevEnv
}