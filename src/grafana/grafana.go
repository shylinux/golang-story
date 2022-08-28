package grafana

import (
	"path"
	"strings"

	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/mdb"
	"shylinux.com/x/icebergs/base/web"
	"shylinux.com/x/icebergs/core/chat/oauth"
	kit "shylinux.com/x/toolkits"
)

type server struct {
	ice.Code
	linux string `data:"https://dl.grafana.com/oss/release/grafana-7.3.4.linux-amd64.tar.gz"`
	list  string `name:"list port path auto start install" help:"可视化"`
}

func (s server) Start(m *ice.Message, arg ...string) {
	s.Code.Start(m, "", "bin/grafana-server", func(p string) {
		port := path.Base(p)
		url := m.Option(ice.MSG_USERWEB)
		hostname := web.OptionUserWeb(m.Message).Hostname()
		client_id := m.Cmdx(oauth.Prefix(oauth.AUTHORIZE), mdb.CREATE, oauth.REDIRECT_URI, kit.Format("http://%s:%s/login/generic_oauth", hostname, port))
		value := kit.KeyValue(nil, "", kit.Dict(
			"server", kit.Dict("domain", hostname, "http_addr", hostname, "http_port", port),
			"security", kit.Dict(
			// "admin_user", m.Option(ice.MSG_USERNAME),
			),
			"auth", kit.Dict(
				// "oauth_auto_login", "true",
				// "disable_login_form", "true",
				// "disable_signout_menu", "true",
				"generic_oauth", kit.Dict(
					"enabled", "true", "client_id", client_id, "scopes", "openid profile email",
					"auth_url", kit.MergeURL2(url, "/chat/oauth/authorize"),
					"token_url", kit.MergeURL2(url, "/chat/oauth/token"),
					"api_url", kit.MergeURL2(url, "/chat/oauth/userinfo"),
				),
			),
		))

		section := ""
		kit.Rewrite(path.Join(p, "conf/defaults.ini"), func(p string) string {
			if strings.TrimSpace(p) == "" {
				return p
			}
			if strings.HasPrefix(p, "#") {
				return p
			}
			if strings.HasPrefix(p, "[") {
				section = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(p), "["), "]"))
				return p
			}
			ls := kit.Split(p, " =", " =", " =")
			if v, ok := value[kit.Keys(section, ls[0])]; ok {
				return kit.Format("%s = %s", ls[0], v)
			}
			return p
		})
	})
}
func (s server) List(m *ice.Message, arg ...string) {
	s.Code.List(m, "", arg...)
}

func init() { ice.CodeCtxCmd(server{}) }
