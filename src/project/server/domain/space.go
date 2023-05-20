package domain

type Space struct {
	Common
	Parent int
	Name   string
}

type SpaceCommon struct {
	Common
	SpaceID int
}
