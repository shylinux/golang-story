module shylinux.com/x/golang-story

go 1.13

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	shylinux.com/x/ice v1.3.3
	shylinux.com/x/icebergs v1.5.11
	shylinux.com/x/toolkits v0.7.6
)

require (
	shylinux.com/x/mysql-story v0.5.9
	shylinux.com/x/redis-story v0.6.1
)
