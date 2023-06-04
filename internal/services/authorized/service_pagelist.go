package authorized

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Page              int64  `json:"page"`               // 第几页
	PageSize          int64  `json:"page_size"`          // 每页显示条数
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (authorizedList []*model.Authorized, err error) {

	offset := (searchData.Page - 1) * searchData.PageSize

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

	authorizedList, err = iAuthorizedDo.
		Limit(int(searchData.PageSize)).
		Offset(int(offset)).
		Order(query.Authorized.ID.Desc()).
		Find()
	if err != nil {
		return nil, err
	}

	return
}
