module shylinux.com/x/golang-story

go 1.15

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.15
	shylinux.com/x/icebergs v1.6.2
	shylinux.com/x/toolkits v0.7.12
)

require (
	github.com/emersion/go-imap v1.2.1
	github.com/emersion/go-message v0.15.0
	github.com/yuin/goldmark v1.5.5
)
