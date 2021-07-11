package infra

import (
	"api/domain/model"
	"api/domain/repository"
	"api/gen/entv1"
	"context"
	"fmt"
	"log"
)

type StatusRepository struct {
	Client *entv1.Client
}

func NewStatusRepository(c *entv1.Client) repository.IStatus {
	return &StatusRepository{
		Client: c,
	}
}

func (c *StatusRepository) Create(status *model.Status, ctx context.Context) (*model.Status, error) {
	entStatus := entv1.Status{
		Title: status.Title.ToString(),
	}
	DBStatus, err := c.Client.Status.
		Create().
		SetTitle(entStatus.Title).
		Save(ctx)
	if err != nil {
		log.Println("failed creating status: %w", err)
		return nil, fmt.Errorf("failed creating status: %w", err)
	}
	log.Println("status was created: ", DBStatus.String())

	return status, nil
}
