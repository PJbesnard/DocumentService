package usecase

import (
	"github.com/PJBesnard/DocumentService/internal/repository"
)

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

func (d *DocumentUseCase) Create(content *repository.Content) (*repository.Document, error) {
	return d.repository.Create(content)
}
