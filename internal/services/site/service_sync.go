/**
 * @Author: chentong
 * @Date: 2024/05/19 下午7:16
 */

package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *service) Sync(ctx core.Context, id int64) error {
	site, err := query.Site.WithContext(ctx.RequestContext()).Where(query.Site.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("站点信息不存在!")
		}
		return err
	}

	if _, err := query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(id)).
		Updates(model.Site{
			Title:       getWebTitle(site),
			Thumb:       getWebLogoIconUrlByUrl(site),
			Description: getWebDescription(site),
			IsUsed:      -1,
		}); err != nil {
		return err
	}

	return nil
}
