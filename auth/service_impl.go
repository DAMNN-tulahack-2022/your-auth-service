package auth

import (
	"context"

	"github.com/damnn/tulahack/your-auth-service/repository"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (acs *AuthService) GithubLogin(ctx context.Context, login string) (id uint32, err error) {
	ghViews, err := scrapGithubViews(login)
	if err != nil {
		return
	}

	user := new(repository.User)
/* 	technologiesMap := make(map[string]struct{}, len(ghViews))
	for _, view := range ghViews {
		technologiesMap[view.Language] = struct{}{}
	}

	for tech := range technologiesMap {
		user.SkillsIds = append(user.SkillsIds, uint32(len(tech)))
	} */

	user.Login = login
	user.AvaterURL = ghViews[0].Owner.URL
	err = acs.repo.FlowUser(ctx, user)
	return user.Id, err
}
