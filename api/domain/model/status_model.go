package model

import "errors"

type StatusID int
type StatusTitle string

type Status struct {
	ID    StatusID
	Title StatusTitle
}

func NewStatusID(id int) (*StatusID, error) {
	if id < 1 {
		err := errors.New("idは1以上の整数にしてください。")
		return nil, err
	}

	statusID := StatusID(id)
	return &statusID, nil
}

func NewTitle(title string) (*StatusTitle, error) {
	if len(title) > 255 {
		err := errors.New("タイトル255字以下にしてください。")
		return nil, err
	}
	statusTitle := StatusTitle(title)
	return &statusTitle, nil
}
