/**
 * @Author: chentong
 * @Date: 2025/01/18 21:59
 */

package index

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) About(ctx *gin.Context) (*v1.AboutResp, error) {
	sysConfig, err := s.configRepo.WithContext(ctx).FindOne()
	if err != nil {
		return nil, err
	}

	return &v1.AboutResp{
		About: v1.About{
			AboutSite:   sysConfig.AboutSite,
			AboutAuthor: sysConfig.AboutAuthor,
			IsAbout:     sysConfig.IsAbout,
		},
	}, nil
}
