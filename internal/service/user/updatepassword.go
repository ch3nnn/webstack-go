/**
 * @Author: chentong
 * @Date: 2024/11/11 18:40
 */

package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/middleware"
)

func (s *service) UpdatePassword(ctx *gin.Context, req *v1.UpdatePasswordReq) (*v1.UpdatePasswordResp, error) {
	user, err := s.userRepo.WithContext(ctx).
		FindOne(
			s.userRepo.WhereByID(ctx.GetInt(middleware.UserID)),
			s.userRepo.WhereByPassword(req.OldPassword),
		)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(errors.New("用户ID"), v1.ErrNotFound)
		}
		return nil, err
	}

	_, err = s.userRepo.WithContext(ctx).Update(&model.SysUser{Password: req.NewPassword}, s.userRepo.WhereByID(user.ID))
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePasswordResp{}, nil
}
