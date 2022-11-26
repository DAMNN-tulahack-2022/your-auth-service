package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	Owner struct {
		URL string `json:"avatar_url"`
	}

	GithubRepoView struct {
		Language string `json:"language"`
		Owner    Owner  `json:"owner"`
	}
)

const (
	_githubUsersUrl = "https://api.github.com/users/"
	_reposPostfix   = "/repos"
)

func scrapGithubViews(githubLogin string) ([]GithubRepoView, error) {
	body, err := http.DefaultClient.Get(fmt.Sprintf(_githubUsersUrl+"%s"+_reposPostfix, githubLogin))
	if err != nil {
		return nil, err
	}

	defer body.Body.Close()
	b, err := io.ReadAll(body.Body)
	if err != nil {
		return nil, err
	}

	views := make([]GithubRepoView, 0)
	if err = json.Unmarshal(b, &views); err != nil {
		return nil, err
	}

	return views, err
}
