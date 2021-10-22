package usecase

import "Vade/documentService/internal/repository"

type DocumentUseCase struct {
	repository repository.DocumentStock
}

func NewDocumentUseCase(repository *repository.DocumentStock) *DocumentUseCase {
	return &DocumentUseCase{repository: *repository}
}

func (d *DocumentUseCase) GetById(id string) (*repository.Document, error) {
	return d.repository.GetById(id)
}

func (d *DocumentUseCase) DeleteById(id string) error {
	return d.repository.DeleteById(id)
}

func (d *DocumentUseCase) UpdateById(document *repository.Document, id string) (*repository.Document, error) {
	return d.repository.UpdateById(document, id)
}
