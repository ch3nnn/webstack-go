package install

import (
	"net/http"
	"runtime"

	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/file"

	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) View() core.HandlerFunc {
	type viewResponse struct {
		Config       configs.Config
		MinGoVersion float64
		GoVersion    string
	}

	return func(ctx core.Context) {
		if _, ok := file.IsExists(configs.ProjectInstallMark); ok {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
		}

		obj := new(viewResponse)
		obj.Config = configs.Get()
		obj.MinGoVersion = configs.MinGoVersion
		obj.GoVersion = runtime.Version()
		ctx.HTML("install_view", obj)
	}
}
