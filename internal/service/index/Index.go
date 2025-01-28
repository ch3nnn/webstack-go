/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:52
 */

package index

import (
	"context"
	"sort"

	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

// buildTree 构建树形结构
func buildTree(nodes []*v1.TreeNode, pid int) []*v1.TreeNode {
	var treeNodes []*v1.TreeNode
	for _, node := range nodes {
		if node.Pid == pid {
			node.Child = buildTree(nodes, node.Id)
			treeNodes = append(treeNodes, node)
		}
	}
	return treeNodes
}

// categoryTree 对树形结构按 Sort 字段排序
func categoryTree(nodes []*v1.TreeNode) []*v1.TreeNode {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Sort < nodes[j].Sort
	})

	for _, node := range nodes {
		if len(node.Child) > 0 {
			categoryTree(node.Child)
		}
	}
	return nodes
}

// categorySites 将站点数据归类到分类站点中
func categorySites(sites []*model.StSite, treeNodes []*v1.TreeNode) (data []*v1.CategorySite) {
	for _, node := range treeNodes {
		categorySite := &v1.CategorySite{
			Category: node.Name,
			SiteList: []model.StSite{},
		}

		for _, site := range sites {
			if site.CategoryID == node.Id {
				categorySite.SiteList = append(categorySite.SiteList, *site)
			}
		}

		if len(categorySite.SiteList) > 0 {
			data = append(data, categorySite)
		}

		if len(node.Child) > 0 {
			childCategorySites := categorySites(sites, node.Child)
			data = append(data, childCategorySites...)
		}
	}

	return data
}

// Index 获取首页数据
func (s *service) Index(ctx context.Context) (*v1.IndexResp, error) {
	var (
		g          errgroup.Group
		sysConfig  *model.SysConfig
		sites      []*model.StSite
		categories []*model.StCategory
	)

	g.Go(func() (err error) {
		categories, err = s.categoryRepo.WithContext(ctx).FindAllOrderBySort(query.StCategory.Sort.Abs(), s.categoryRepo.WhereByIsUsed(true))
		return err
	})

	g.Go(func() (err error) {
		sites, err = s.siteRepo.WithContext(ctx).FindAll(s.siteRepo.WhereByIsUsed(true))
		return err
	})

	g.Go(func() (err error) {
		sysConfig, err = s.configRepo.WithContext(ctx).FindOne()
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	nodes := make([]*v1.TreeNode, len(categories))
	for i, category := range categories {
		nodes[i] = &v1.TreeNode{
			Id:   category.ID,
			Pid:  category.ParentID,
			Name: category.Title,
			Icon: category.Icon,
			Sort: category.Sort,
		}
	}

	categoryTree := categoryTree(buildTree(nodes, 0))
	categorySites := categorySites(sites, categoryTree)

	return &v1.IndexResp{
		ConfigSite: &v1.ConfigSite{
			SiteTitle:   sysConfig.SiteTitle,
			SiteKeyword: sysConfig.SiteKeyword,
			SiteDesc:    sysConfig.SiteDesc,
			SiteRecord:  sysConfig.SiteRecord,
			SiteLogo:    sysConfig.SiteLogo,
			SiteFavicon: sysConfig.SiteFavicon,
		},
		About: &v1.About{
			AboutSite:   sysConfig.AboutSite,
			AboutAuthor: sysConfig.AboutAuthor,
			IsAbout:     sysConfig.IsAbout,
		},
		CategoryTree:  categoryTree,
		CategorySites: categorySites,
	}, nil
}
