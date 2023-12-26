package config

import "os"

type Configuration struct {
	ServerPort       string
	ExecutionMode    string
	LocalEnvironment string
}

const (
	ServerPort              = "SERVER_PORT"
	DefaultServerPort       = "8090"
	ExecutionMode           = "EXECUTION_MODE"
	DefaultExecutionMode    = "default"
	LocalEnvironment        = "LOCAL_ENVIRONMENT"
	DefaultLocalEnvironment = "default"
)

var config *Configuration

func GetConfig() *Configuration {
	if config == nil {
		config = initConfig()
	}

	return config
}

func initConfig() *Configuration {
	return &Configuration{
		ServerPort:       getEnvOrDefault(ServerPort, DefaultServerPort),
		ExecutionMode:    getEnvOrDefault(ExecutionMode, DefaultExecutionMode),
		LocalEnvironment: getEnvOrDefault(LocalEnvironment, DefaultLocalEnvironment),
	}
}

func getEnvOrDefault(environmentVarName string, defValue string) string {
	if environmentVar := os.Getenv(environmentVarName); environmentVar != "" {
		return environmentVar
	}

	return defValue
}
