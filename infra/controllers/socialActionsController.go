package controllers

import (
	"encoding/json"
	"go-social-action/application/usecase"
	"go-social-action/infra/http/responses"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type SocialActionController struct {
	*usecase.CreateSocialActionUseCase
	*usecase.FindSocialActionUseCase
	*usecase.FindSocialActionsUseCase
	*usecase.DeleteSocialActionUseCase
	*usecase.UpdateSocialActionUseCase
}

func (c *SocialActionController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	var input usecase.CreateSocialActionInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	socialAction, err := c.CreateSocialActionUseCase.Execute(r.Context(), &input)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.Created(w, socialAction)
}

func (c *SocialActionController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	volunteer, err := c.FindSocialActionUseCase.Execute(r.Context(), id)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.Ok(w, volunteer)
}

func (c *SocialActionController) GetAll(w http.ResponseWriter, r *http.Request) {
	volunteers, err := c.FindSocialActionsUseCase.Execute(r.Context())
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.Ok(w, volunteers)
}

func (c *SocialActionController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := c.DeleteSocialActionUseCase.Execute(r.Context(), id)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.NoContent(w)
}

func (c *SocialActionController) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	var input usecase.UpdateSocialActionInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}
	id := chi.URLParam(r, "id")
	err = c.UpdateSocialActionUseCase.Execute(r.Context(), id, &input)
	if err != nil {
		responses.ResponseWithErr(w, err)
		return
	}
	responses.NoContent(w)
}
