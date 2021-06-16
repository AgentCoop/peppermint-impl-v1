package webproxy

import (
	"github.com/AgentCoop/peppermint/cmd"
	"github.com/AgentCoop/peppermint/internal/plugin"
	g "github.com/AgentCoop/peppermint/internal/plugin/webproxy/grpc/server"
	model "github.com/AgentCoop/peppermint/internal/plugin/webproxy/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	Name = "WebProxy"
)

type webProxy struct {
	plugin.WebProxyConfigurator
}

func init() {
	proxy := &webProxy{
		NewConfigurator(),
	}
	proxy.WebProxyConfigurator = NewConfigurator()
	reg := runtime.GlobalRegistry()
	serviceInfo := &runtime.ServiceInfo{
		Name: Name,
		Cfg: proxy.WebProxyConfigurator,
		Initializer: proxy.initializer,
	}
	reg.RegisterService(serviceInfo)
	reg.RegisterParserCmdHook(cmd.CMD_NAME_DB_MIGRATE, proxy.migrateDb)
}

func (w *webProxy) initializer() runtime.Service {
	proxy := g.NewServer(
		Name,
		w.WebProxyConfigurator.Address(),
		w.WebProxyConfigurator,
		w,
	)
	return proxy
}

func (w *webProxy) migrateDb(options interface{}) {
	db := runtime.GlobalRegistry().Db()
	h := db.Handle()
	h.AutoMigrate(&model.WebProxyConfig{})
}

