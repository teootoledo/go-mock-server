package resources

type CreateMockRequest struct {
	Endpoint   string `json:"endpoint" example:"/api/example" binding:"required"`
	Method     string `json:"method" example:"POST" binding:"required"`
	Payload    string `json:"payload" example:"{\"example\":\"example\"}"`
	StatusCode int    `json:"status-code" example:"200" binding:"required"`
}

type MockCreatedResponse struct {
	Message string `json:"message" example:"Mock created successfully"`
}

type Mock struct {
	Payload    string `json:"payload" example:"{\"example\":\"example\"}"`
	StatusCode int    `json:"status-code" example:"200"`
}

type MockStorage struct {
	Storage map[string]Mock
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		Storage: make(map[string]Mock),
	}
}
