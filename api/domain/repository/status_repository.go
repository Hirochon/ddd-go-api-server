package repository

import "api/domain/model"

type IStatus interface {
	Create(status *model.Status) (*model.Status, error)
}
