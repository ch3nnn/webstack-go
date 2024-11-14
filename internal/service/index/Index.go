/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:52
 */

package index

import (
	"context"

	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

func buildTree(nodes []*v1.TreeNode, pid int) (treeNodes []*v1.TreeNode) {
	for _, node := range nodes {
		if node.Pid == pid {
			node.Child = buildTree(nodes, node.Id)
			treeNodes = append(treeNodes, node)
		}
	}
	return treeNodes
}

func (s *service) Index(ctx context.Context) (*v1.IndexResponseData, error) {
	var (
		g          errgroup.Group
		sites      []*model.StSite
		categories []*model.StCategory
	)

	g.Go(func() (err error) {
		categories, err = s.categoryRepo.WithContext(ctx).FindAllOrderBySort(query.StCategory.Sort.Abs(), s.categoryRepo.WhereByIsUsed(true))
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() (err error) {
		sites, err = s.siteRepo.WithContext(ctx).FindAll(s.siteRepo.WhereByIsUsed(true))
		if err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	nodes := make([]*v1.TreeNode, len(categories))
	for i, category := range categories {
		nodes[i] = &v1.TreeNode{
			Id:    category.ID,
			Pid:   category.ParentID,
			Name:  category.Title,
			Icon:  category.Icon,
			Child: nil,
		}
	}

	categorySites := make([]*v1.CategorySite, len(categories))
	for i, category := range categories {
		var siteList []model.StSite
		for _, st := range sites {
			if category.ID == st.CategoryID {
				siteList = append(siteList, *st)
			}
		}
		categorySites[i] = &v1.CategorySite{
			Category: category.Title,
			SiteList: siteList,
		}
	}

	return &v1.IndexResponseData{
		CategoryTree:  buildTree(nodes, 0),
		CategorySites: categorySites,
	}, nil
}
