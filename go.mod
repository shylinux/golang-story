module shylinux.com/x/golang-story

go 1.15

replace (
	shylinux.com/x/go-git => ./usr/go-git
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.11
	shylinux.com/x/icebergs v1.5.18
	shylinux.com/x/toolkits v0.7.9
)

require (
	github.com/cvilsmeier/sqinn-go v1.1.2 // indirect
	github.com/glebarez/sqlite v1.9.0 // indirect
	github.com/yuin/goldmark v1.4.13
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.2 // indirect
)
