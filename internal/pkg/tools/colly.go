package tools

import (
	"crypto/tls"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net"
	"net/http"
	"time"
)

func NewColly() *colly.Collector {
	collector := colly.NewCollector(
		colly.DetectCharset(),   // 检测响应编码
		colly.IgnoreRobotsTxt(), // 忽略 robots 协议
	)

	// HTTP 设置
	collector.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,  // 超时时间
			KeepAlive: 30 * time.Second, // KeepAlive 超时时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,                                  // 关闭 keepalive
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 不安全的跳过验证
	})

	// 随机 useragent 请求头
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)
	return collector
}
