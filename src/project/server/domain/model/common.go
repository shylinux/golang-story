package model

type Model interface {
	TableName() string
	GetKey() string
	GetID() string
}
