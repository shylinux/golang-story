module shylinux.com/x/golang-story

go 1.15

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.16
	shylinux.com/x/icebergs v1.6.4
	shylinux.com/x/toolkits v0.7.13
)

require (
	github.com/emersion/go-imap v1.2.1
	github.com/emersion/go-message v0.15.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/yuin/goldmark v1.5.5
)
