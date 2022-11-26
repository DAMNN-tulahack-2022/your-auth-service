package transport

import (
	"context"

	"github.com/damnn/tulahack/your-auth-service/gen/proto"
)

func (gp *AuthProxy) GithubLogin(ctx context.Context, request *proto.GithubLoginRequest) (response *proto.GithubLoginResponse, err error)  {
	id, err := gp.service.GithubLogin(ctx, request.GetLogin())
	if err != nil {
		return
	} 

	return &proto.GithubLoginResponse{Id: id, Login: request.Login}, nil
}




