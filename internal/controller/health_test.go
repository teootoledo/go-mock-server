package controller_test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/baloo.v3"
	"io"
	"mock-server/internal/app"
	"mock-server/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HealthControllerSuite struct {
	suite.Suite
}

func TestHealthControllerSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerSuite))
}

func (suite *HealthControllerSuite) TestGetPing() {
	gin.SetMode(gin.TestMode)
	testApp := app.NewApp()
	testApp.Setup()

	healthController := controller.NewHealth()
	testApp.HealthController = healthController

	// Gin attaches dependencies by value so after changing dependencias at app lvl it is required to re-generate gin's engine
	testApp.ConfigureRoutes()

	testServer := httptest.NewServer(testApp.Engine)

	request := baloo.New(testServer.URL).
		Get("/v1/ping").
		Expect(suite.T())

	_ = request.
		Status(http.StatusOK).
		AssertFunc(assertHealthResponse(controller.HealthResponse{Status: "pong"})).
		Done()
}

func assertHealthResponse(responseBody controller.HealthResponse) func(res *http.Response, req *http.Request) error {
	function := func(res *http.Response, req *http.Request) error {
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		var result controller.HealthResponse
		_ = json.Unmarshal(body, &result)

		if !cmp.Equal(responseBody, result) {
			return errors.New("error")
		}

		return nil
	}

	return function
}
