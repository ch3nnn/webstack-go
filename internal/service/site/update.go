/**
 * @Author: chentong
 * @Date: 2024/06/30 下午10:14
 */

package site

import (
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func getIconBase64ByFormFile(req *v1.SiteUpdateReq) string {
	file, err := req.File.Open()
	if err != nil {
		return defaultIcon
	}
	defer file.Close()

	imgData, err := io.ReadAll(file)
	if err != nil {
		return defaultIcon
	}

	return base64.StdEncoding.EncodeToString(imgData)
}

func getIconBase64ByURL(req *v1.SiteUpdateReq) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get(req.Icon)
	if err != nil {
		return defaultIcon
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultIcon
	}

	return base64.StdEncoding.EncodeToString(imgData)
}

func (s *service) Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error) {
	update := make(map[string]any)

	if req.CategoryId != 0 {
		update["CategoryID"] = req.CategoryId
	}
	if req.Title != "" {
		update["Title"] = req.Title
	}
	if req.Icon != "" {
		update["Icon"] = getIconBase64ByURL(req)
	}
	if req.File != nil {
		update["Icon"] = getIconBase64ByFormFile(req)
	}
	if req.Description != "" {
		update["Description"] = req.Description
	}
	if req.Url != "" {
		update["Url"] = req.Url
	}
	if req.IsUsed != nil {
		update["IsUsed"] = req.IsUsed
	}
	if req.Sort >= 0 {
		update["Sort"] = req.Sort
	}

	_, err = s.siteRepository.WithContext(ctx).Update(update, s.siteRepository.WhereByID(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.SiteUpdateResp{ID: req.Id}, nil
}
