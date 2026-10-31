package tools

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
  <circle cx="12" cy="12" r="10" fill="#0071e3"/>
</svg>`

func pngBytes(t *testing.T, w, h int) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{R: 30, G: 120, B: 220, A: 255})
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatalf("encode png: %v", err)
	}
	return buf.Bytes()
}

/** 把字节内容包成 multipart.FileHeader，模拟表单上传 */
func multipartFileHeader(t *testing.T, name string, content []byte) *multipart.FileHeader {
	t.Helper()
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	fw, err := w.CreateFormFile("file", name)
	if err != nil {
		t.Fatalf("create form file: %v", err)
	}
	if _, err := fw.Write(content); err != nil {
		t.Fatalf("write form file: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	if err := req.ParseMultipartForm(1 << 20); err != nil {
		t.Fatalf("parse multipart: %v", err)
	}
	return req.MultipartForm.File["file"][0]
}

func TestMultipartImgToBase64_SVG原样保留(t *testing.T) {
	header := multipartFileHeader(t, "logo.svg", []byte(testSVG))

	got, err := MultipartImgToBase64(header, 64, 64)
	assert.NoError(t, err)

	// SVG 是矢量图：不能像位图那样重采样，必须原样存下，只补 data URI 前缀
	assert.True(t, strings.HasPrefix(got, "data:image/svg+xml;base64,"), "前缀不对: %s", got)

	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(got, "data:image/svg+xml;base64,"))
	assert.NoError(t, err)
	assert.Equal(t, testSVG, string(raw), "SVG 内容被改动了")
}

func TestMultipartImgToBase64_SVG带XML声明(t *testing.T) {
	svg := `<?xml version="1.0" encoding="UTF-8"?>` + testSVG
	header := multipartFileHeader(t, "logo.svg", []byte(svg))

	got, err := MultipartImgToBase64(header, 64, 64)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(got, "data:image/svg+xml;base64,"), "带 XML 声明的 SVG 没被识别: %s", got)
}

func TestMultipartImgToBase64_位图仍然缩放(t *testing.T) {
	header := multipartFileHeader(t, "logo.png", pngBytes(t, 200, 50))

	got, err := MultipartImgToBase64(header, 64, 64)
	assert.NoError(t, err)

	// 位图保持老行为：纯 base64（无 data: 前缀），并缩放到目标尺寸
	assert.False(t, strings.HasPrefix(got, "data:"), "位图不该带 data: 前缀")

	raw, err := base64.StdEncoding.DecodeString(got)
	assert.NoError(t, err)
	img, err := png.Decode(bytes.NewReader(raw))
	assert.NoError(t, err)
	assert.Equal(t, 64, img.Bounds().Dx())
	assert.Equal(t, 64, img.Bounds().Dy())
}

func TestURLImgToBase64_SVG(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(testSVG))
	}))
	defer srv.Close()

	got, err := URLImgToBase64(srv.URL, 64, 64)
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(got, "data:image/svg+xml;base64,"), "外链 SVG 没被识别: %s", got)
}

func TestImgSrc(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"空值", "", ""},
		{"外链", "https://example.com/a.svg", "https://example.com/a.svg"},
		{"站内相对路径", "/assets/logo.svg", "/assets/logo.svg"},
		{"SVG data URI 原样返回", "data:image/svg+xml;base64,PHN2Zy8+", "data:image/svg+xml;base64,PHN2Zy8+"},
		{"位图裸 base64 补 png 前缀", "iVBORw0KGgo=", "data:image/png;base64,iVBORw0KGgo="},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, ImgSrc(c.in))
		})
	}
}

func TestMultipartImgToBase64_非SVG的XML不算SVG(t *testing.T) {
	// XHTML 错误页 / RSS 也以 <?xml 开头，不能被当成 SVG 存成图片
	xhtml := `<?xml version="1.0"?><html><body>404 Not Found</body></html>`
	header := multipartFileHeader(t, "page.xml", []byte(xhtml))

	_, err := MultipartImgToBase64(header, 64, 64)
	assert.Error(t, err, "非 SVG 的 XML 应当按位图解码失败报错，而不是被存成 SVG")
}

func TestMultipartImgToBase64_SVG带DOCTYPE和注释(t *testing.T) {
	cases := map[string]string{
		"DOCTYPE": `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` + testSVG,
		"注释":      "<!-- generator: figma -->\n" + testSVG,
		"空白+声明":   "  \n<?xml version=\"1.0\"?>\n" + testSVG,
	}
	for name, content := range cases {
		t.Run(name, func(t *testing.T) {
			header := multipartFileHeader(t, "logo.svg", []byte(content))
			got, err := MultipartImgToBase64(header, 64, 64)
			assert.NoError(t, err)
			assert.True(t, strings.HasPrefix(got, "data:image/svg+xml;base64,"), "没被识别为 SVG: %s", got)
		})
	}
}

func TestMultipartImgToBase64_超大SVG报错(t *testing.T) {
	big := testSVG + strings.Repeat("<rect/>", 8*1024) // ~40KB
	assert.Greater(t, len(big), maxSVGBytes)
	header := multipartFileHeader(t, "big.svg", []byte(big))

	_, err := MultipartImgToBase64(header, 64, 64)
	assert.Error(t, err, "超过上限的 SVG 应当报错，而不是原样入库把首页撑大")
}
