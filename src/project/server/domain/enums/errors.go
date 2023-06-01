package enums

var Errors = struct {
	Success       int
	Unknown       int
	InvalidParams int
	ModelCreate   int
	ModelRemove   int
	ModelInfo     int
	ModelList     int
}{
	Success:       0,
	Unknown:       10000,
	InvalidParams: 10001,
	ModelCreate:   20001,
	ModelRemove:   20002,
	ModelInfo:     20003,
	ModelList:     20004,
}
