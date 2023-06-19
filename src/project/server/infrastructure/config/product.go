package config

type Product struct {
	Name   string
	Repos  string
	Portal []Portal
}
type Portal struct {
	Name     string
	Title    string
	Redirect string
	Display  string
	Views    []Views
}
type Views struct {
	Name    string
	Icon    string
	Title   string
	Service string
	View    []View
}
type View struct {
	Name    string
	Icon    string
	Title   string
	Source  string
	Service string
	Display string
}
