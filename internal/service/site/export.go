/**
 * @Author: chentong
 * @Date: 2025/01/29 19:10
 */

package site

import (
	"strconv"

	"github.com/gin-gonic/gin"
	excelize "github.com/xuri/excelize/v2"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
)

var (
	sheetName = "Sheet1"
	headers   = []string{"ID", "Logo", "名称简介", "链接", "分类", "创建日期", "更新日期", "状态"}
)

func (s *service) Export(ctx *gin.Context, req *v1.SiteExportReq) (resp *v1.SiteExportResp, err error) {
	var siteCategories []repository.SiteCategory
	_, err = s.siteRepository.WithContext(ctx).FindSiteCategoryWithPage(1, 10000, &siteCategories, s.siteRepository.LikeInByTitleOrDescOrURL(req.Search))
	if err != nil {
		return nil, err
	}

	excelFile := excelize.NewFile()

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := excelFile.SetCellValue(sheetName, cell, header); err != nil {
			continue
		}
	}

	for i, siteCategory := range siteCategories {
		row := strconv.Itoa(i + 2)

		excelFile.SetCellValue(sheetName, "A"+row, siteCategory.StSite.ID)
		excelFile.SetCellValue(sheetName, "B"+row, siteCategory.StSite.Icon)
		excelFile.SetCellValue(sheetName, "C"+row, siteCategory.StSite.Title)
		excelFile.SetCellValue(sheetName, "D"+row, siteCategory.StSite.URL)
		excelFile.SetCellValue(sheetName, "E"+row, siteCategory.StCategory.Title)
		excelFile.SetCellValue(sheetName, "F"+row, siteCategory.StSite.CreatedAt)
		excelFile.SetCellValue(sheetName, "G"+row, siteCategory.StSite.UpdatedAt)
		excelFile.SetCellValue(sheetName, "H"+row, siteCategory.StSite.IsUsed)
	}

	index, err := excelFile.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	excelFile.SetActiveSheet(index)

	return &v1.SiteExportResp{File: excelFile}, nil
}
