package authorized

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) List(ctx core.Context, searchData *SearchData) (authorizedList []*model.Authorized, err error) {

	iAuthorizedDo := query.Authorized.WithContext(ctx.RequestContext())
	iAuthorizedDo = iAuthorizedDo.Where(query.Authorized.IsDeleted.Eq(-1))
	if searchData.BusinessKey != "" {
		iAuthorizedDo = iAuthorizedDo.Where(query.Authorized.BusinessKey.Eq(searchData.BusinessKey))
	}

	if searchData.BusinessSecret != "" {
		iAuthorizedDo = iAuthorizedDo.Where(query.Authorized.BusinessSecret.Eq(searchData.BusinessSecret))
	}

	if searchData.BusinessDeveloper != "" {
		iAuthorizedDo = iAuthorizedDo.Where(query.Authorized.BusinessDeveloper.Eq(searchData.BusinessDeveloper))
	}
	authorizedList, err = iAuthorizedDo.Order(query.Authorized.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	return
}
