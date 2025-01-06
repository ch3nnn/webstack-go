/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:56
 */

package site

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/mat/besticon/besticon"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

// Service 层
var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error)
	Delete(ctx context.Context, req *v1.SiteDeleteReq) (resp *v1.SiteDeleteResp, err error)
	Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error)
	BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (resp *v1.SiteCreateResp, err error)
	Sync(ctx *gin.Context, req *v1.SiteSyncReq) (resp *v1.SiteSyncResp, err error)
}

type service struct {
	*s.Service
	siteRepository     repository.IStSiteDao
	categoryRepository repository.IStCategoryDao
}

func NewService(s *s.Service) Service {
	return &service{
		Service:            s,
		siteRepository:     repository.NewStSiteDao(),
		categoryRepository: repository.NewStCategoryDao(),
	}
}

func (s *service) i() {}

// -----------------------------------------------------------------------------------------------------------------------------------------

// default icon base64
const defaultIcon = "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAMAAACdt4HsAAAA+VBMVEUAAAC6sWUzNjZCQkIsLi8sLjAsLi//tSs6Oj8tLzAsLi/9sycsLi/9sycsLi/9tCctLi/8tSj/tSguMTEuMzQzMzMwMzb/ujP9tCcsLi8uLzEtLy8sLi8sLi8tLi8sLi8tLi8sLi8tLjAsLi8sLy8sLi8sLjAsLjAtLy8tLzAtLzAuMDAuMDL/tyn9tCf8tCctLi/8tCf8tCf8tCf9sycsLjD9tCgsLjD9tSf9tCgtLzAtMDAuLzAvLzEvLzL/ty0tLy/+tij9tCj9tCj9syf9syj9tCf9tCj9tCf9tCn8tSj/tCj/tikwMDT/vzP/v0AsLy/8sycsLi+FeN01AAAAUXRSTlMAAg8F/uOqIwv7+Pjp6dLQdWM/KSIcFg3v7Uv18tzZzce5ta+ej4eBb2lUQT0x+8jBvrWvqpqLiollZGBcNjAck1T039fNo5h7cFpHOB8UCKLRW+BFAAACmElEQVRYw82VaVPyMBSFW5aC7FulgMiqICIIIiq4gPvum/7/H/NOe+00xwZL+OTzhZkm5wz35txE+eOo2ZExyqob63NpZpHObSaPTJjDJCIvL57HmUv8vChbfJIhyawqVzwg2YrICRNzEpEoHpBohXoBxXtIXqhSxUu0QlS8ZCuK0zhbj/i0KFu8fyuO0kyO9BHqDSaLAQ5jJs+YNzCYPAZvMGTyDHmDVybPq8KTjUvK41kFuZlI6Sc3gjl4WFv+kPO7hpD1L6fidOArH8AgeIiM/eITWSUtdIP279v9L/L7N3tPsFv4KY/ux8xwj1rxYqzK3gsV3wubsf0o6AN10+L4WrH4fBqIin/6tFevj+299QBv0DaJ0Ok2TefIox/R9G2fhkyizRuETYfSdyuWdyC/W34XXzIdwrzBrumyW7Y/fT3fMofb5y/7Wxn28QZnJk8zT+l+ZMQj5TbfhF1nvEFQg7WtPerQ+9Aa23fq894W7NGCCjBPwHJiplKwlhQcdfZjfa6IksBT63OL/RqsuSkACi0T0D4U4kPDhRaXQ+QwBRsrnYBVfKcCX1OHv72tByW0aC1aKC8dqCuniZq9wGYjW4uAuxfp10JtinG+sVLfyFOc2yGnxZiknQydba8qlFd7lJnMDuTIJuPsKbt7ANe97LhneIOU4F+CXFBfijdIeGN8Vef09StvnBO8gSY6qcuwM7iXojPWeAO9KspKtBOzctuJilJW1RVwEKdV7/7r6uKc6z4xhnmBSYM4Y4ydVkLl0A0iIYgzNBl7T+eBxyRCdGuFmrPyrBny3ndi4N4E4Mb1AWIsirMv9HYA+Or4Q68XQu+eBL0wqunllSEKN2GF4iyHrkFuN6FfwxdCHnXeKDXmqvKn+Q+oeE3vIQF62QAAAABJRU5ErkJggg=="

func getWebLogoIconBase64(url string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	b := besticon.New(
		besticon.WithLogger(besticon.NewDefaultLogger(io.Discard)), // disable verbose logging
		besticon.WithHTTPClient(client),
	)

	icons, err := b.NewIconFinder().FetchIcons(url)
	if err != nil || len(icons) == 0 {
		return defaultIcon, err
	}

	resp, err := client.Get(icons[0].URL)
	if err != nil {
		return defaultIcon, err
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return defaultIcon, err
	}

	return base64.StdEncoding.EncodeToString(imgData), nil
}

func getWebTitle(url string) (title string, err error) {
	c := tools.NewColly()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title += e.Text
	})
	if err := c.Visit(url); err != nil {
		return "", err
	}

	return
}

func getWebDescription(url string) (doc string, err error) {
	c := tools.NewColly()
	c.OnXML("//meta[@name='description']/@content|//meta[@name='Description']/@content|//meta[@name='DESCRIPTION']",
		func(e *colly.XMLElement) {
			doc += e.Text
		},
	)
	if err := c.Visit(url); err != nil {
		return "", err
	}

	return
}
