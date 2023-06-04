package authorized

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
}

func (s *service) ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (authorizedAPIS []*model.AuthorizedAPI, err error) {
	iAuthorizedAPIDo := query.AuthorizedAPI.WithContext(ctx.RequestContext())
	iAuthorizedAPIDo = iAuthorizedAPIDo.Where(query.AuthorizedAPI.IsDeleted.Eq(-1))

	if searchAPIData.BusinessKey != "" {
		iAuthorizedAPIDo = iAuthorizedAPIDo.Where(query.AuthorizedAPI.BusinessKey.Eq(searchAPIData.BusinessKey))
	}
	authorizedAPIS, err = iAuthorizedAPIDo.Order(query.AuthorizedAPI.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	return
}
