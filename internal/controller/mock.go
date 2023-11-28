package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mock-server/internal/constants"
	"mock-server/internal/external/resources"
	"mock-server/internal/logs"
	"net/http"
)

var responseMap = make(map[string]map[string]resources.Mock)
var bannedEndpoints = []string{
	"/v1/mock",
	"/v1/docs",
	"/v1/docs/index.html",
}

type Mock interface {
	SetMockResponse(ctx *gin.Context)
	NotFound(ctx *gin.Context)
}

type mockController struct {
	logger logs.Logger
}

// SetMockResponse godoc
// @Summary Set mock response
// @Description Set mock response
// @Tags mock
// @Accept json
// @Produce json
// @Param			body		body		resources.CreateMockRequest		true	"Mock creation details"
// @Success 200 {object} string "Mock response set successfully"
// @Failure 400 {object} string "Invalid JSON"
// @Router /mock [post]
func (mc *mockController) SetMockResponse(ctx *gin.Context) {
	var request resources.CreateMockRequest
	if err := ctx.BindJSON(&request); err != nil {
		mc.logger.Error(ctx, "error in MockController#SetMockResponse: bad request", err)
		SetErrorResponseBadRequest(ctx, http.StatusBadRequest, constants.BadRequestMessage, err.Error())
		return
	}

	if err := mc.evaluateCreateMockRequest(request); err != nil {
		mc.logger.Error(ctx, "error in MockController#SetMockResponse: ", err)
		SetErrorResponseBadRequest(ctx, http.StatusBadRequest, constants.BadRequestMessage, err.Error())
		return
	}

	if _, exists := responseMap[request.Endpoint]; !exists {
		responseMap[request.Endpoint] = make(map[string]resources.Mock)
	}

	responseMap[request.Endpoint][request.Method] = resources.Mock{
		Payload:    request.Payload,
		StatusCode: request.StatusCode,
	}

	ctx.JSON(http.StatusOK, resources.MockCreatedResponse{
		Message: "Mock response set successfully",
	})
}

// NotFound godoc
// @Summary Not found
// @Description Matches the request path with the response map and returns the response if found
// @Tags mock
// @Accept json
// @Produce json
// @Success 200 {object} object "Mock response"
// @Failure 404 {object} string "Not Found"
func (mc *mockController) NotFound(ctx *gin.Context) {
	// Obtain the endpoint from the request path
	endpoint := ctx.Request.URL.Path
	method := ctx.Request.Method

	// Check if the endpoint exists in the response map
	if response, isInMap := responseMap[endpoint][method]; isInMap {
		// If the endpoint exists, return the response
		ctx.JSON(response.StatusCode, response.Payload)
	} else {
		// If the endpoint does not exist, return a 404
		mc.logger.Error(ctx, "error in MockController#NotFound: mock not found for the given endpoint and method", nil)
		SetErrorResponseNotFound(ctx, http.StatusNotFound, constants.NotFoundMessage, "Mock not found for the given endpoint and method")
	}
}

// PRIVATE METHODS =================================================================================
func (mc *mockController) evaluateCreateMockRequest(request resources.CreateMockRequest) error {
	for _, bannedEndpoint := range bannedEndpoints {
		if request.Endpoint == bannedEndpoint {
			return errors.New("endpoint is not allowed")
		}
	}

	if request.Method != http.MethodGet && request.Method != http.MethodPost && request.Method != http.MethodPut && request.Method != http.MethodDelete {
		return errors.New("method is invalid")
	}

	if request.StatusCode == 0 || request.StatusCode < 100 || request.StatusCode > 599 {
		return errors.New("status code is required and must be between 100 and 599")
	}

	return nil
}

// NewMock - returns a new mock controller
func NewMock() Mock {
	logger := logs.New("Mock Controller")

	return &mockController{
		logger: logger,
	}
}
