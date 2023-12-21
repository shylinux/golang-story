module shylinux.com/x/golang-story

go 1.15

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.4.2
	shylinux.com/x/icebergs v1.8.2
	shylinux.com/x/toolkits v0.8.2
)

require (
	github.com/emersion/go-imap v1.2.1
	github.com/emersion/go-message v0.15.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/yuin/goldmark v1.5.5
)
