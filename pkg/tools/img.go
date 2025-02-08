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

	"github.com/disintegration/imaging"
)

func resizeImg2Base64(r io.Reader, width, height int) (base64Str string, err error) {
	img, err := imaging.Decode(r, imaging.AutoOrientation(true))
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

// ResizeMultipartImgToBase64 将multipart.FileHeader表示的图片文件调整大小，并以base64编码字符串的形式返回。
// 参数f是包含图片文件信息的multipart.FileHeader指针；
// 参数width和height分别是目标图片的宽度和高度。
// 返回值base64Str是调整大小后的图片的base64编码字符串；err是错误信息，如果执行过程中发生错误则不为nil。
func ResizeMultipartImgToBase64(f *multipart.FileHeader, width, height int) (base64Str string, err error) {
	file, err := f.Open()
	if err != nil {
		return
	}
	defer file.Close()

	return resizeImg2Base64(file, width, height)
}

// ResizeURLImgToBase64 从指定的URL获取图像，并将其调整为指定的宽度和高度后，转换为Base64编码的字符串。
// 参数:
//
//	url - 图像的URL地址。
//	width - 调整后图像的宽度。
//	height - 调整后图像的高度。
//
// 返回值:
//
//	base64Str - 转换后的Base64编码字符串。
//	err - 错误信息，如果执行过程中遇到任何错误，则返回该错误。
func ResizeURLImgToBase64(url string, width, height int) (base64Str string, err error) {
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

	return resizeImg2Base64(resp.Body, width, height)
}
