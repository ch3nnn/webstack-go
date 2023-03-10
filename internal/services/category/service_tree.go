package category

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/category"
)

type TreeNode struct {

	// Id 节点ID
	Id int32
	// Pid 父节点ID
	Pid int32
	// Name 节点名称
	Name string
	// Icon 图标
	Icon string
	// Child 获取子节点切片
	Child []*TreeNode
}

func (s *service) Tree(ctx core.Context) (nodes []*TreeNode, err error) {
	qb := category.NewQueryBuilder()
	listData, err := qb.
		WhereIsUsed(mysql.EqualPredicate, 1).
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	for _, v := range listData {
		nodes = append(nodes, &TreeNode{
			Id:    v.Id,
			Pid:   v.ParentId,
			Name:  v.Title,
			Icon:  v.Icon,
			Child: nil,
		})
	}
	treeNode := buildTree(nodes, 0)
	return treeNode, nil
}

func buildTree(nodes []*TreeNode, pid int32) []*TreeNode {
	res := make([]*TreeNode, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			v.Child = buildTree(nodes, v.Id)
			res = append(res, v)
		}
	}
	return res
}
