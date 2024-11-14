/**
 * @Author: chentong
 * @Date: 2024/05/26 上午12:27
 */

package user

import (
	"context"
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Login(ctx context.Context, req *v1.LoginReq) (resp *v1.LoginResp, err error) {
	user, err := s.userRepo.WithContext(ctx).
		FindOne(
			s.userRepo.WhereByUsername(req.Username),
			s.userRepo.WhereByPassword(req.Password),
		)
	if err != nil {
		return nil, v1.ErrorUserNameAndPassword
	}

	token, err := s.Jwt.GenToken(user.ID, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, v1.ErrorTokenGeneration
	}

	return &v1.LoginResp{Token: token}, nil
}
