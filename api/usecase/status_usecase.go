package usecase

import (
	"api/domain/model"
	"api/domain/repository"
)

type StatusUsecase interface {
	Create(id int, title string) (*model.Status, error)
}

type statusUsecase struct {
	statusRepo repository.IStatus
}

func NewStatusUsecase(statusRepo repository.IStatus) StatusUsecase {
	return &statusUsecase{
		statusRepo: statusRepo,
	}
}

func (su *statusUsecase) Create(id int, title string) (*model.Status, error) {
	statusID, err := model.NewStatusID(id)
	if err != nil {
		return nil, err
	}
	statusTitle, err := model.NewStatusTitle(title)
	if err != nil {
		return nil, err
	}
	return &model.Status{
		ID:    *statusID,
		Title: *statusTitle,
	}, nil
}
