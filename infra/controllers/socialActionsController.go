package controllers

import (
	"encoding/json"
	"go-social-action/application/usecase"
	"go-social-action/infra/http/responses"
	"io/ioutil"
	"net/http"
)

type SocialActionController struct {
	*usecase.CreateSocialActionUseCase
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
