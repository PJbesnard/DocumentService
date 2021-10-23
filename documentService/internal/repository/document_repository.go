package repository

import (
	"github.com/PJBesnard/DocumentService/documentService/internal/entity"
	"github.com/PJBesnard/DocumentService/documentService/internal/utils"
)

type Content struct {
	Name        string `json:"name" binding:"required`
	Description string `json:"description"`
}

type Document struct {
	DocumentId string  `json:"document_id"`
	Content    Content `json:"content"`
}

type DocumentStock struct {
	Documents map[string]Document `json:"documents"`
}

func NewDocumentStock(documents map[string]Document) *DocumentStock {
	return &DocumentStock{Documents: documents}
}

func (d *DocumentStock) GetById(id string) (*Document, error) {
	if val, ok := d.Documents[id]; ok {
		return &val, nil
	} else {
		return nil, entity.ErrDocumentNotFound
	}
}

func (d *DocumentStock) DeleteById(id string) error {
	if _, ok := d.Documents[id]; ok {
		delete(d.Documents, id)
		return nil
	} else {
		return entity.ErrDocumentNotFound
	}
}

func (d *DocumentStock) Create(content *Content) (*Document, error) {
	doc := &Document{
		DocumentId: utils.RandomId(24),
		Content:    *content,
	}
	d.Documents[doc.DocumentId] = *doc
	return doc, nil
}
