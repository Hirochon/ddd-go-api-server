package handler

import (
	"api/gen/restapiv1/operations/hoge"
	"api/usecase"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type StatusHandler interface {
	POST() func(http.ResponseWriter, *http.Request)
}

type statusHandler struct {
	statusUsecase usecase.StatusUsecase
}

func NewStatusHandler(statusUsecase usecase.StatusUsecase) StatusHandler {
	return &statusHandler{
		statusUsecase: statusUsecase,
	}
}

func (sh *statusHandler) POST() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		var hoge hoge.PostIDStatusBody
		err := hoge.UnmarshalJSON(buf.Bytes())
		if err != nil {
			fmt.Fprintf(w, "Bad Request! you can't unmaeshal %v\n", err)
			return
		}
		// fmt.Fprintf(w, "もらった: %v\n", hoge.Title)
		sh.statusUsecase.Create(int(hoge.Object.ID), hoge.Title, r.Context())
		w.WriteHeader(http.StatusCreated)
	}
}
