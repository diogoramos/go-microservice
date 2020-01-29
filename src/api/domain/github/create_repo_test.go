package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T){
	request := CreateRepoRequest{
		Name:        "Go",
		Description: "Lang",
		Homepage:    "https://uol.com.br",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	//Receive an interface and try to create a valid json
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.Equal(t, `{"name":"Go","description":"Lang","homepage":"https://uol.com.br","private":false,"has_issues":false,"has_projects":false,"has_wiki":false}`, string(bytes))
	//fmt.Println(string(bytes))
}