module shylinux.com/x/golang-story

go 1.11

require (
	shylinux.com/x/ice v1.0.8
	shylinux.com/x/icebergs v1.3.3
	shylinux.com/x/toolkits v0.6.7
)

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)