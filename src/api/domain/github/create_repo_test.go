package github_test

import (
	"encoding/json"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := github.CreateRepoRequest{
		Name:        "Golang Test",
		Description: "A golang test",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target github.CreateRepoRequest

	// Unmarshal takes an input byte array and a pointer that we're trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}
