package repository

import "Vade/documentService/internal/entity"

type Document struct {
	DocumentId  string `json:"document_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
		delete(d.Documents, "id")
		return nil
	} else {
		return entity.ErrDocumentNotFound
	}
}

func (d *DocumentStock) UpdateById(document *Document, id string) (*Document, error) {
	if val, ok := d.Documents[id]; ok {
		d.Documents[id] = *document
		return &val, nil
	} else {
		return nil, entity.ErrDocumentNotFound
	}
}
