package assets

import "embed"

var (
	//go:embed static
	Static embed.FS

	//go:embed templates
	Templates embed.FS

	//go:embed admin/dist
	SPA embed.FS
)
