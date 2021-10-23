package usecase

import (
	"Vade/documentService/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentUseCase_GetById(t *testing.T) {
	documents := make(map[string]repository.Document)
	t.Run("success", func(t *testing.T) {
		documents["testId"] = repository.Document{DocumentId: "testId", Content: *&repository.Content{Description: "Description test", Name: "Name test"}}
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		res, err := usecase.GetById("testId")
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("error-failed", func(t *testing.T) {
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		res, err := usecase.GetById("testId")
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestDocumentUseCase_DeleteById(t *testing.T) {
	documents := make(map[string]repository.Document)
	t.Run("success", func(t *testing.T) {
		documents["testId"] = repository.Document{DocumentId: "testId", Content: *&repository.Content{Description: "Description test", Name: "Name test"}}
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		err := usecase.DeleteById("testId")
		assert.NoError(t, err)
	})
	t.Run("error-failed", func(t *testing.T) {
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		err := usecase.DeleteById("testId")
		assert.Error(t, err)
	})
}

func TestDocumentUseCase_Create(t *testing.T) {
	documents := make(map[string]repository.Document)
	t.Run("success", func(t *testing.T) {
		documents["testId"] = repository.Document{DocumentId: "testId", Content: *&repository.Content{Description: "Description test", Name: "Name test"}}
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		err := usecase.DeleteById("testId")
		assert.NoError(t, err)
	})
	t.Run("error-failed", func(t *testing.T) {
		repository := repository.NewDocumentStock(documents)
		usecase := NewDocumentUseCase(repository)
		err := usecase.DeleteById("testId")
		assert.Error(t, err)
	})
}
