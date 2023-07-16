module shylinux.com/x/golang-story

go 1.15

// replace (
// 	shylinux.com/x/ice => ./usr/release
// 	shylinux.com/x/icebergs => ./usr/icebergs
// 	shylinux.com/x/toolkits => ./usr/toolkits
// )

require (
	shylinux.com/x/ice v1.3.11
	shylinux.com/x/icebergs v1.5.16
	shylinux.com/x/toolkits v0.7.8
)

require (
	google.golang.org/grpc v1.55.0
	shylinux.com/x/mysql-story v0.5.10
	shylinux.com/x/redis-story v0.6.2
)
