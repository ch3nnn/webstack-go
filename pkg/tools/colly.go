package tools

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func NewColly() *colly.Collector {
	collector := colly.NewCollector(
		colly.DetectCharset(),   // 检测响应编码
		colly.IgnoreRobotsTxt(), // 忽略 robots 协议
	)

	// HTTP 设置
	collector.WithTransport(&http.Transport{
		MaxIdleConns:          100,             // 最大空闲连接数
		IdleConnTimeout:       5 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   5 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,                                  // 关闭 keepalive
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 不安全的跳过验证
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,  // 超时时间
			KeepAlive: 30 * time.Second, // KeepAlive 超时时间
		}).DialContext,
	})

	// 随机 user agent 请求头
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)
	return collector
}
