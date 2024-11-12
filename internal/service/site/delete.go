/**
 * @Author: chentong
 * @Date: 2024/06/10 上午12:20
 */

package site

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Delete(ctx context.Context, req *v1.SiteDeleteReq) (resp *v1.SiteDeleteResp, err error) {
	err = s.siteRepository.WithContext(ctx).Delete(s.siteRepository.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	return &v1.SiteDeleteResp{ID: req.ID}, nil
}
