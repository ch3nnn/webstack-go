/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:48
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Delete(ctx context.Context, req *v1.CategoryDeleteReq) (*v1.CategoryDeleteResp, error) {
	return nil, s.categoryRepo.WithContext(ctx).Delete(s.categoryRepo.WhereByID(req.ID))
}
