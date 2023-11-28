package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/baloo.v3"
	"mock-server/internal/app"
	"mock-server/internal/controller"
	"mock-server/internal/external/resources"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockControllerSuite struct {
	suite.Suite
}

func TestMockControllerSuite(t *testing.T) {
	suite.Run(t, new(MockControllerSuite))
}

func (suite *MockControllerSuite) SetupTest() {
	suite.T().Setenv("EXECUTION_MODE", "test")
	suite.T().Setenv("LOCAL_ENVIRONMENT", "true")
}

type setMockResponseCase struct {
	testName           string
	mockRequest        resources.CreateMockRequest
	expectedResponse   resources.MockCreatedResponse
	expectedStatusCode int
	assert             func(res *http.Response, req *http.Request) error
}

func (suite *MockControllerSuite) TestSetMockResponse() {
	testCases := []setMockResponseCase{
		{
			testName: "Mock response is set successfully",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/health",
				Method:     "POST",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 200,
			},
			expectedResponse: resources.MockCreatedResponse{
				Message: "Mock response created successfully",
			},
			expectedStatusCode: http.StatusOK,
			assert:             controller.AssertSuccess,
		},
		{
			testName: "Mock response is set successfully without payload",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/health",
				Method:     "POST",
				StatusCode: 200,
			},
			expectedResponse: resources.MockCreatedResponse{
				Message: "Mock response created successfully",
			},
			expectedStatusCode: http.StatusOK,
			assert:             controller.AssertSuccess,
		},
		{
			testName: "Unable to set mock response due to empty endpoint",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "",
				Method:     "POST",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 200,
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Unable to set mock response due to invalid endpoint",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/v1/mock",
				Method:     "POST",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 200,
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Unable to set mock response due to empty method",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/health",
				Method:     "",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 200,
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Unable to set mock response due to invalid method",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/health",
				Method:     "invalid",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 200,
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Unable to set mock response due to invalid status code",
			mockRequest: resources.CreateMockRequest{
				Endpoint:   "/health",
				Method:     "POST",
				Payload:    "{\"status\": \"ok\"}",
				StatusCode: 999,
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Unable to set mock response due to missing status code",
			mockRequest: resources.CreateMockRequest{
				Endpoint: "/health",
				Method:   "POST",
				Payload:  "{\"status\": \"ok\"}",
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
	}

	for _, test := range testCases {
		suite.T().Run(test.testName, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			testApp := app.NewApp()
			testApp.Setup()

			mockController := controller.NewMock()
			testApp.MockController = mockController

			// Gin attaches dependencies by value so after changing dependencias at app lvl it is required to re-generate gin's engine
			testApp.ConfigureRoutes()

			testServer := httptest.NewServer(testApp.Engine)
			request := baloo.New(testServer.URL).
				Post("/v1/mock").
				JSON(test.mockRequest).
				Expect(t)

			_ = request.
				Status(test.expectedStatusCode).
				AssertFunc(test.assert).
				Done()
		})
	}
}

type getMockResponseCase struct {
	testName           string
	endpoint           string
	method             string
	expectedStatusCode int
	assert             func(res *http.Response, req *http.Request) error
}

func (suite *MockControllerSuite) TestGetMockResponse() {

	testCases := []getMockResponseCase{
		{
			testName:           "Mock response is returned successfully",
			endpoint:           "/health",
			method:             "POST",
			expectedStatusCode: http.StatusOK,
			assert:             controller.AssertSuccess,
		},
		{
			testName:           "Mock response is not found",
			endpoint:           "/health",
			method:             "PUT",
			expectedStatusCode: http.StatusNotFound,
			assert:             controller.AssertNotFound,
		},
	}

	for _, test := range testCases {
		suite.T().Run(test.testName, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			testApp := app.NewApp()
			testApp.Setup()

			mockController := controller.NewMock()
			testApp.MockController = mockController

			// Gin attaches dependencies by value so after changing dependencias at app lvl it is required to re-generate gin's engine
			testApp.ConfigureRoutes()

			testServer := httptest.NewServer(testApp.Engine)

			// Set mock response
			_ = baloo.New(testServer.URL).
				Post("/v1/mock").
				JSON(resources.CreateMockRequest{
					Endpoint:   "/health",
					Method:     "POST",
					Payload:    "{\"status\": \"ok\"}",
					StatusCode: 200,
				}).
				Expect(t).
				Status(200).
				Done()

			// Now depending on the test.method execute the test request
			var request *baloo.Expect
			switch test.method {
			case "POST":
				request = baloo.New(testServer.URL).
					Post(test.endpoint).Expect(t)
			case "PUT":
				request = baloo.New(testServer.URL).
					Put(test.endpoint).Expect(t)
			}

			_ = request.
				Status(test.expectedStatusCode).
				AssertFunc(test.assert).
				Done()
		})
	}
}
