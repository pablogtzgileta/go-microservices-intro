package services

import (
	"github.com/pablogtzgileta/go-microservices-intro/src/api/config"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/domain/github"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/domain/repositories"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/providers/github_provider"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/utils/errors"
	"strings"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}
