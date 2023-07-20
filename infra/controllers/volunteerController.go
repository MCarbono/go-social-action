package controllers

import (
	"encoding/json"
	"go-social-action/application/usecase"
	"go-social-action/infra/http/responses"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type VolunteerController struct {
	*usecase.CreateVolunteerUseCase
	*usecase.FindVolunteerUseCase
}

func (c *VolunteerController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	var createVolunteerInput usecase.CreateVolunteerInput
	err = json.Unmarshal(body, &createVolunteerInput)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	volunteer, err := c.CreateVolunteerUseCase.Execute(r.Context(), &createVolunteerInput)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.Created(w, volunteer)
}

func (c *VolunteerController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	volunteer, err := c.FindVolunteerUseCase.Execute(r.Context(), id)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.Ok(w, volunteer)
}
