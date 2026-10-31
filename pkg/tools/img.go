/**
 * @Author: chentong
 * @Date: 2025/02/08 12:37
 */

package tools

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
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
	// SVG 是原样入库的矢量图：超过这个尺寸会撑大首页，也可能撞上数据库 text 字段上限
	maxSVGBytes = 32 << 10
	// 判类型只看开头这么多字节
	sniffBytes = 1024
)

// isSVG 判断内容是不是 SVG：跳过 BOM、空白、XML 声明/PI/DOCTYPE/注释后，看是否以 <svg 开头。
// 不能只认扩展名/Content-Type —— 上传的文件头不可信，外链也常给 application/octet-stream；
// 也不能见到 <?xml 就当 SVG —— XHTML 错误页和 RSS 都以它开头。
func isSVG(data []byte) bool {
	if len(data) > sniffBytes {
		data = data[:sniffBytes]
	}
	return bytes.HasPrefix(skipXMLProlog(data), []byte("<svg"))
}

// skipXMLProlog 跳过 BOM、空白，以及 <??> / <!-- --> / <!DOCTYPE ...> 这些前导声明。
func skipXMLProlog(data []byte) []byte {
	for {
		data = bytes.TrimLeft(bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF}), " \t\r\n")

		var end int
		switch {
		case bytes.HasPrefix(data, []byte("<?")):
			end = bytes.Index(data, []byte("?>"))
			end += 2
		case bytes.HasPrefix(data, []byte("<!--")):
			end = bytes.Index(data, []byte("-->"))
			end += 3
		case len(data) >= 9 && bytes.EqualFold(data[:9], []byte("<!doctype")):
			end = bytes.IndexByte(data, '>')
			end++
		default:
			return data
		}

		if end <= 0 { // 声明被截断了，交给后面的匹配失败
			return data
		}
		data = data[end:]
	}
}

// img2Base64 把图片字节转成可直接入库的字符串。
// SVG 是矢量图，重采样只会丢精度，所以原样保留并补上 data URI 前缀；
// 位图仍按 width/height 缩放，返回纯 base64（历史数据就是这个格式）。
func img2Base64(data []byte, width, height int) (base64Str string, err error) {
	if isSVG(data) {
		if len(data) > maxSVGBytes {
			return "", fmt.Errorf("SVG 体积 %d 字节，超过上限 %d 字节", len(data), maxSVGBytes)
		}
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

	data, err := io.ReadAll(io.LimitReader(resp.Body, maxRemoteImageBytes+1))
	if err != nil {
		return
	}
	if len(data) > maxRemoteImageBytes {
		return "", fmt.Errorf("外链图片超过上限 %d 字节", maxRemoteImageBytes)
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
