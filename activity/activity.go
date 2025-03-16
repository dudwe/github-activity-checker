package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GitResponse []struct {
	CommitType string `json:"type"`
	Actor      struct {
	} `json:"actor"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		RepositoryID int `json:"repository_id"`
		Commits      []struct {
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

func GetActivityRequest(urlString string) (*http.Response, error) {

	parsedUrl, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	res, err := http.Get(parsedUrl.String())
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	fmt.Printf("client:got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	return res, nil
}

func GetActivity(user_name string) (string, error) {
	request_string := "https://api.github.com/users/" + user_name + "/events"
	res, err := GetActivityRequest(request_string)
	if err != nil {
		return "", fmt.Errorf("error making GetActivity request: %w", err)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("request reponse abnormal: %d", res.StatusCode)
	}

	var gitReponse GitResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&gitReponse)
	if err != nil {
		return "", fmt.Errorf("failed to decode json: %w", err)
	}
	if len(gitReponse) == 0 {
		return "No Commits for " + user_name, nil
	}
	resultString := ""
	for _, gitCommit := range gitReponse {
		commitType := gitCommit.CommitType
		repoName := gitCommit.Repo.Name
		gitPayload := gitCommit.Payload
		repoId := gitPayload.RepositoryID
		for _, gitCommitPayload := range gitPayload.Commits {
			resultString += fmt.Sprintf(" %s %s %d: %s:%s\n", commitType, repoName, repoId, gitCommitPayload.Author.Name, gitCommitPayload.Message)
		}
	}
	fmt.Print(resultString)
	return resultString, nil
}
