package assets

import "embed"

var (
	//go:embed static
	Static embed.FS

	//go:embed templates
	Templates embed.FS

	// all: 前缀必须保留——Vite 产物里有以 _ 开头的 chunk，默认 embed 规则会跳过它们
	//go:embed all:admin/dist
	SPA embed.FS
)
