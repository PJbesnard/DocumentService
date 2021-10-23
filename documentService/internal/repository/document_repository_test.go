package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentStock_GetById(t *testing.T) {
	documents := make(map[string]Document)
	t.Run("success", func(t *testing.T) {
		documents["testId"] = Document{DocumentId: "testId", Content: *&Content{Description: "Description test", Name: "Name test"}}
		repository := NewDocumentStock(documents)
		res, err := repository.GetById("testId")
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("error-failed", func(t *testing.T) {
		repository := NewDocumentStock(documents)
		res, err := repository.GetById("testId")
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestDocumentStock_Create(t *testing.T) {
	documents := make(map[string]Document)
	repository := NewDocumentStock(documents)
	res, err := repository.Create(&Content{Description: "Description test", Name: "Name test"})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	res2, err := repository.GetById(res.DocumentId)
	assert.NoError(t, err)
	assert.NotNil(t, res2)
}

func TestDocumentStock_DeleteById(t *testing.T) {
	documents := make(map[string]Document)
	t.Run("success", func(t *testing.T) {
		documents["testId"] = Document{DocumentId: "testId", Content: *&Content{Description: "Description test", Name: "Name test"}}
		repository := NewDocumentStock(documents)
		err := repository.DeleteById("testId")
		assert.NoError(t, err)
	})
	t.Run("error-failed", func(t *testing.T) {
		repository := NewDocumentStock(documents)
		err := repository.DeleteById("testId")
		assert.Error(t, err)
	})
}
