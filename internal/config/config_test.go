package config_test

import (
	"mock-server/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) SetupTest() {
	suite.T().Setenv("EXECUTION_MODE", "")
	suite.T().Setenv("LOCAL_ENVIRONMENT", "")
}

func (suite *ConfigSuite) TestCustomValuesAreReturnedProperly() {
	conf := config.GetConfig()

	require.Equal(suite.T(), conf.ServerPort, config.DefaultServerPort)
	require.Equal(suite.T(), conf.ExecutionMode, config.DefaultExecutionMode)
	require.Equal(suite.T(), conf.LocalEnvironment, config.DefaultLocalEnvironment)
}
