package handler

import (
	"encoding/json"
	"github.com/PJBesnard/DocumentService/documentService/internal/repository"
	"github.com/PJBesnard/DocumentService/documentService/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DocumentServer struct {
	documentUseCase *usecase.DocumentUseCase
}

func NewDocumentServer(documentUseCase *usecase.DocumentUseCase) *DocumentServer {
	return &DocumentServer{documentUseCase: documentUseCase}
}

func (d *DocumentServer) GetDocumentById() func(*gin.Context) {
	return func(c *gin.Context) {
		doc, err := d.documentUseCase.GetById(c.Param("id"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		res, err := json.Marshal(doc)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, string(res))
	}
}

func (d *DocumentServer) CreateDocument() func(*gin.Context) {
	return func(c *gin.Context) {
		var content repository.Content
		err := c.BindJSON(&content)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		doc, err := d.documentUseCase.Create(&content)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		res, err := json.Marshal(doc)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, string(res))
	}
}

func (d *DocumentServer) DeleteDocumentById() func(*gin.Context) {
	return func(c *gin.Context) {
		err := d.documentUseCase.DeleteById(c.Param("id"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusOK, "")
	}
}

func DocumentHandler(router *gin.Engine) {
	documents := make(map[string]repository.Document)
	documentRepo := repository.NewDocumentStock(documents)
	documentUsecase := usecase.NewDocumentUseCase(documentRepo)
	documentServer := NewDocumentServer(documentUsecase)

	router.GET("/document/:id", documentServer.GetDocumentById())
	router.POST("/document", documentServer.CreateDocument())
	router.DELETE("/document/:id", documentServer.DeleteDocumentById())
}
