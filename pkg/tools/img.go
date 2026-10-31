/**
 * @Author: chentong
 * @Date: 2025/02/08 12:37
 */

package tools

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	svgDataURIPrefix = "data:image/svg+xml;base64,"
	pngDataURIPrefix = "data:image/png;base64,"
	// 外链图片读入内存的上限，避免被超大响应拖垮
	maxRemoteImageBytes = 8 << 20
)

// isSVG 判断内容是不是 SVG：跳过 UTF-8 BOM 和空白后，看是否以 <svg 或 <?xml 开头。
// 不能只认扩展名/Content-Type —— 上传的文件头不可信，外链也常给 application/octet-stream。
func isSVG(data []byte) bool {
	trimmed := bytes.TrimLeft(bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF}), " \t\r\n")
	return bytes.HasPrefix(trimmed, []byte("<svg")) || bytes.HasPrefix(trimmed, []byte("<?xml"))
}

// img2Base64 把图片字节转成可直接入库的字符串。
// SVG 是矢量图，重采样只会丢精度，所以原样保留并补上 data URI 前缀；
// 位图仍按 width/height 缩放，返回纯 base64（历史数据就是这个格式）。
func img2Base64(data []byte, width, height int) (base64Str string, err error) {
	if isSVG(data) {
		return svgDataURIPrefix + base64.StdEncoding.EncodeToString(data), nil
	}

	img, err := imaging.Decode(bytes.NewReader(data), imaging.AutoOrientation(true))
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

// MultipartImgToBase64 把表单上传的图片转成入库字符串（SVG 原样保留，位图缩放到 width×height）。
func MultipartImgToBase64(f *multipart.FileHeader, width, height int) (base64Str string, err error) {
	file, err := f.Open()
	if err != nil {
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return
	}

	return img2Base64(data, width, height)
}

// URLImgToBase64 抓取外链图片并转成入库字符串（SVG 原样保留，位图缩放到 width×height）。
// 参数 url 是图片地址，width/height 是位图的目标尺寸。
func URLImgToBase64(url string, width, height int) (base64Str string, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, maxRemoteImageBytes))
	if err != nil {
		return
	}

	return img2Base64(data, width, height)
}

// ImgSrc 把库里的图标值变成可以直接放进 <img src> / <link href> 的值。
// 库里可能是：外链、本站相对路径、SVG 的 data URI，或位图的纯 base64。
func ImgSrc(icon string) string {
	if icon == "" {
		return ""
	}
	if strings.HasPrefix(icon, "data:") ||
		strings.HasPrefix(icon, "http") ||
		strings.HasPrefix(icon, "/") {
		return icon
	}
	return pngDataURIPrefix + icon
}
