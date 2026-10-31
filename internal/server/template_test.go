package server

import (
	"bytes"
	"html"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func renderImg(t *testing.T, icon string) string {
	t.Helper()
	tmpl := template.Must(template.New("t").Funcs(templateFuncs()).Parse(`<img src="{{ imgSrc . }}">`))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, icon); err != nil {
		t.Fatalf("执行模板: %v", err)
	}
	return buf.String()
}

// 断言浏览器最终看到的 src：html/template 会把 + 转成 &#43;，解码后仍是原来的值
func assertImgSrc(t *testing.T, icon, want string) {
	t.Helper()
	got := renderImg(t, icon)
	// html/template 的 URL 过滤器默认把 data: 整个换成 #ZgotmplZ，图片会变空白
	assert.NotContains(t, got, "ZgotmplZ")
	assert.Equal(t, want, html.UnescapeString(got))
}

func TestImgSrc模板函数_data协议不被过滤(t *testing.T) {
	assertImgSrc(t, "data:image/svg+xml;base64,PHN2Zz48L3N2Zz4=",
		`<img src="data:image/svg+xml;base64,PHN2Zz48L3N2Zz4=">`)
	assertImgSrc(t, "data:image/png;base64,a+b",
		`<img src="data:image/png;base64,a+b">`)
}

func TestImgSrc模板函数_位图裸base64补全前缀(t *testing.T) {
	assertImgSrc(t, "iVBORw0KGgo=", `<img src="data:image/png;base64,iVBORw0KGgo=">`)
}

func TestImgSrc模板函数_外链原样输出(t *testing.T) {
	assertImgSrc(t, "https://example.com/a.svg", `<img src="https://example.com/a.svg">`)
}

func TestImgSrc模板函数_引号不会逸出属性(t *testing.T) {
	got := renderImg(t, `x" onerror="alert(1)`)
	assert.NotContains(t, got, `onerror="alert(1)">`, "属性被闭合了：%s", got)
}
