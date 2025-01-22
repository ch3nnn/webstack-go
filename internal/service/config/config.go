/**
 * @Author: chentong
 * @Date: 2025/01/18 21:10
 */

package config

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) GetConfig(ctx context.Context) (*v1.ConfigResp, error) {
	conf, err := s.configRepo.WithContext(ctx).FindOne()
	if err != nil {
		return nil, err
	}

	return &v1.ConfigResp{
		ID:          conf.ID,
		AboutSite:   conf.AboutSite,
		AboutAuthor: conf.AboutAuthor,
		IsAbout:     conf.IsAbout,
		SiteTitle:   conf.SiteTitle,
		SiteKeyword: conf.SiteKeyword,
		SiteDesc:    conf.SiteDesc,
		SiteRecord:  conf.SiteRecord,
	}, nil
}
