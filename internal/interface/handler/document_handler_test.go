package handler

import (
	"github.com/PJBesnard/DocumentService/internal/repository"
	"github.com/PJBesnard/DocumentService/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDocumentServer_GetDocumentById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		documents := make(map[string]repository.Document)
		documents["testId"] = repository.Document{DocumentId: "testId", Content: *&repository.Content{Description: "Description test", Name: "Name test"}}
		repository := repository.NewDocumentStock(documents)
		usecase := usecase.NewDocumentUseCase(repository)
		server := NewDocumentServer(usecase)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/document/:id", server.GetDocumentById())
		req, _ := http.NewRequest("GET", "/document/testId", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("error-failed", func(t *testing.T) {
		documents := make(map[string]repository.Document)
		repository := repository.NewDocumentStock(documents)
		usecase := usecase.NewDocumentUseCase(repository)
		server := NewDocumentServer(usecase)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/document/:id", server.GetDocumentById())
		req, _ := http.NewRequest("GET", "/document/testId", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
}

func TestDocumentServer_DeleteDocumentById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		documents := make(map[string]repository.Document)
		documents["testId"] = repository.Document{DocumentId: "testId", Content: *&repository.Content{Description: "Description test", Name: "Name test"}}
		repository := repository.NewDocumentStock(documents)
		usecase := usecase.NewDocumentUseCase(repository)
		server := NewDocumentServer(usecase)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/document/:id", server.DeleteDocumentById())
		req, _ := http.NewRequest("DELETE", "/document/testId", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("error-failed", func(t *testing.T) {
		documents := make(map[string]repository.Document)
		repository := repository.NewDocumentStock(documents)
		usecase := usecase.NewDocumentUseCase(repository)
		server := NewDocumentServer(usecase)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/document/:id", server.DeleteDocumentById())
		req, _ := http.NewRequest("GET", "/document/testId", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)
	})
}

func TestDocumentServer_CreateDocument(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		documents := make(map[string]repository.Document)
		repository := repository.NewDocumentStock(documents)
		usecase := usecase.NewDocumentUseCase(repository)
		server := NewDocumentServer(usecase)
		w := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/document", server.CreateDocument())
		req, _ := http.NewRequest("POST", "/document",
			strings.NewReader(`{
			"description" : "This is a document",
			"name" : "Super Document"
			}`))
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}
