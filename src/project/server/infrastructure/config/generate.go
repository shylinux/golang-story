package config

import "flag"

type Generate struct {
	Path   string
	PbPath string
	TsPath string
	GoPath string
	ShPath string
	JsPath string
}

func init() {
	flag.StringVar(&config.Generate.Path, "generate.path", "idl/", "")
	flag.StringVar(&config.Generate.PbPath, "generate.pbpath", "idl/pb/", "")
	flag.StringVar(&config.Generate.TsPath, "generate.tspath", "idl/ts/", "")
	flag.StringVar(&config.Generate.ShPath, "generate.shpath", "idl/cli/", "")
	flag.StringVar(&config.Generate.GoPath, "generate.gopath", "idl/api/", "")
	flag.StringVar(&config.Generate.JsPath, "generate.jspath", "usr/vue-element-admin/src/api/", "")
}
