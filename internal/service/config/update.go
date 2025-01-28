/**
 * @Author: chentong
 * @Date: 2025/01/18 14:21
 */

package config

import (
	"bytes"
	"encoding/base64"
	"mime/multipart"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
)

const (
	LogoWidth     = 200
	LogoHeight    = 50
	FaviconWidth  = 32
	FaviconHeight = 32
)

func resize2Image(f *multipart.FileHeader, width, height int) (base64Str string, err error) {
	file, err := f.Open()
	if err != nil {
		return
	}
	defer file.Close()

	img, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		return
	}

	var buf bytes.Buffer
	resize := imaging.Resize(img, width, height, imaging.Lanczos)
	if err = imaging.Encode(&buf, resize, imaging.PNG); err != nil {
		return
	}

	base64Str = base64.StdEncoding.EncodeToString(buf.Bytes())

	return
}

func (s *service) Update(ctx *gin.Context, req *v1.ConfigUpdateReq) (resp *v1.ConfigUpdateResp, err error) {
	update := make(map[string]any)
	if req.SiteTitle != nil {
		update["site_title"] = *req.SiteTitle
	}
	if req.SiteDesc != nil {
		update["site_desc"] = *req.SiteDesc
	}
	if req.SiteKeyword != nil {
		update["site_keyword"] = *req.SiteKeyword
	}
	if req.SiteRecord != nil {
		update["site_record"] = *req.SiteRecord
	}
	if req.AboutSite != nil {
		update["about_site"] = *req.AboutSite
	}
	if req.AboutAuthor != nil {
		update["about_author"] = *req.AboutAuthor
	}
	if req.IsAbout != nil {
		update["is_about"] = *req.IsAbout
	}
	if req.LogFile != nil {
		base64Str, err := resize2Image(req.LogFile, LogoWidth, LogoHeight)
		if err != nil {
			base64Str = repository.DefaultLogoBase64
		}

		update["site_logo"] = base64Str
	}
	if req.FaviconFile != nil {
		base64Str, err := resize2Image(req.FaviconFile, FaviconWidth, FaviconHeight)
		if err != nil {
			base64Str = repository.DefaultFaviconBase64
		}

		update["site_favicon"] = base64Str
	}

	if _, err = s.configRepo.WithContext(ctx).Update(update, s.configRepo.WhereByID(1)); err != nil {
		return nil, err
	}

	return nil, nil
}
