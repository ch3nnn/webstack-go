package index

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/category"
	"github.com/ch3nnn/webstack-go/internal/services/site"
	"net/http"
)

type indexResponse struct {
	CategoryTree  []*category.TreeNode
	CategorySites []*site.CategorySite
}

func (h *handler) Index() core.HandlerFunc {
	return func(c core.Context) {

		categoryTree, err := h.categoryService.Tree(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}

		categorySites, err := h.siteService.CategorySite(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}

		response := indexResponse{
			CategoryTree:  categoryTree,
			CategorySites: categorySites,
		}

		c.HTML("index", response)
	}
}
