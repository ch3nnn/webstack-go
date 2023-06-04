package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type TreeNode struct {
	// Id 节点ID
	Id int64
	// Pid 父节点ID
	Pid int64
	// Name 节点名称
	Name string
	// Icon 图标
	Icon string
	// Child 获取子节点切片
	Child []*TreeNode
}

func buildTree(nodes []*TreeNode, pid int64) []*TreeNode {
	res := make([]*TreeNode, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			v.Child = buildTree(nodes, v.Id)
			res = append(res, v)
		}
	}
	return res
}

func (s *service) Tree(ctx core.Context) (nodes []*TreeNode, err error) {
	categories, err := query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.IsUsed.Eq(1)).
		Order(query.Category.Sort).Find()
	if err != nil {
		return nil, err
	}
	for _, cat := range categories {
		nodes = append(nodes, &TreeNode{
			Id:    cat.ID,
			Pid:   cat.ParentID,
			Name:  cat.Title,
			Icon:  cat.Icon,
			Child: nil,
		})
	}
	return buildTree(nodes, 0), nil
}
