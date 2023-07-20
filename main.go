package main

import (
	"fmt"
	"go-social-action/application/usecase"
	"go-social-action/configs"
	"go-social-action/idGenerator"
	"go-social-action/infra/controllers"
	"go-social-action/infra/database"
	"go-social-action/infra/http/router"
	"go-social-action/infra/repository"
	"log"
	"net/http"
)

func main() {
	cfg, err := configs.LoadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Open(cfg.PSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	idGenerator := idGenerator.New()
	volunteerRepository := repository.NewVolunteerRepositoryPostgres(db)
	socialActionRepository := repository.NewSocialActionRepositoryPostgres(db)
	volunteerController := controllers.VolunteerController{
		CreateVolunteerUseCase: usecase.NewCreateVolunteerUseCase(volunteerRepository, idGenerator),
		FindVolunteerUseCase:   usecase.NewFindVolunteerUseCase(volunteerRepository),
	}
	socialActionController := controllers.SocialActionController{
		CreateSocialActionUseCase: usecase.NewCreateSocialActionUseCase(volunteerRepository, socialActionRepository, idGenerator),
	}
	r := router.New(volunteerController, socialActionController)
	fmt.Printf("Starting the server on port %v\n", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
