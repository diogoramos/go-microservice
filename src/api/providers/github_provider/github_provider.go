package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/diogoramos/go-microservice/src/api/domain/github"
	"github.com/diogoramos/go-microservice/src/api/clients/restclient"
	"io/ioutil"

	"net/http"
)

var (
	headerAuthorization = "abc"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = ""
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprint(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, r github.CreateRepoRequest )(*github.CreateRepoResponse, *github.GitHubErrorResponse){
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	//Doing the rest request
	response, err := restclient.Post(urlCreateRepo, r, headers)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode:      http.StatusInternalServerError,
			Message:         err.Error(),
		}
	}

	//Reading the json body
	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}

	//defer is used for execute a function when the main function is returning for the caller.
	//when you use defer you're creating a stack of instructions to be removed/deleted
	defer response.Body.Close()

	//validating the response by statuscode
	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err:= json.Unmarshal(bytes, &result); err != nil{
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid json response body",
		}
	}

	return &result, nil
}