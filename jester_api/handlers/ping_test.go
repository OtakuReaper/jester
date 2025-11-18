package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type fakePing struct {}

func (f *fakePing) Message() string {
	return "fake pong"
}

func TestPingHandler(test *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	handler := NewPingHandler(&fakePing{})
	router.GET("/ping", handler.Ping)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		test.Fatalf("Expected 200, got %d", writer.Code)
	}

	expected := `{"message":"fake pong"}`
	if writer.Body.String() != expected {
		test.Fatalf("Expected %s, got %s", expected, writer.Body.String())
	}
}