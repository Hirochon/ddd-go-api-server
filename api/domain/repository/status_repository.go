package repository

import (
	"api/domain/model"
	"context"
)

type IStatus interface {
	Create(status *model.Status, ctx context.Context) (*model.Status, error)
}
