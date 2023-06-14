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
	flag.StringVar(&config.Matrix.Generate.Path, "generate.path", "idl/", "")
	flag.StringVar(&config.Matrix.Generate.PbPath, "generate.pbpath", "idl/pb/", "")
	flag.StringVar(&config.Matrix.Generate.TsPath, "generate.tspath", "idl/ts/", "")
	flag.StringVar(&config.Matrix.Generate.ShPath, "generate.shpath", "idl/cli/", "")
	flag.StringVar(&config.Matrix.Generate.GoPath, "generate.gopath", "idl/api/", "")
	flag.StringVar(&config.Matrix.Generate.JsPath, "generate.jspath", "usr/vue-element-admin/src/api/", "")
}
