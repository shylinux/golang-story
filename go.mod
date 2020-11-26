module github.com/shylinux/golang-story

go 1.13

require (
	github.com/shylinux/icebergs v0.3.0
	github.com/shylinux/toolkits v0.2.0
)

replace (
	github.com/shylinux/icebergs => ./usr/icebergs
	github.com/shylinux/toolkits => ./usr/toolkits
)
