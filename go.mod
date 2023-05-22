module shylinux.com/x/golang-story

go 1.13

replace (
	shylinux.com/x/ice => ./usr/release
	shylinux.com/x/icebergs => ./usr/icebergs
	shylinux.com/x/mysql-story => ./usr/mysql-story
	shylinux.com/x/redis-story => ./usr/redis-story
	shylinux.com/x/toolkits => ./usr/toolkits
)

require (
	google.golang.org/grpc v1.55.0 // indirect
	shylinux.com/x/ice v1.3.2
	shylinux.com/x/icebergs v1.5.6
	shylinux.com/x/mysql-story v0.5.8
	shylinux.com/x/redis-story v0.6.0
	shylinux.com/x/toolkits v0.7.5
)
