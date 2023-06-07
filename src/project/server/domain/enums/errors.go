package enums

var Errors = struct {
	Success       int64
	Unknown       int64
	NotFoundProxy int64
	InvalidParams int64
	ModelCreate   int64
	ModelRemove   int64
	ModelInfo     int64
	ModelList     int64
}{
	Success:       0,
	Unknown:       10000,
	NotFoundProxy: 10001,
	InvalidParams: 10002,
	ModelCreate:   20001,
	ModelRemove:   20002,
	ModelInfo:     20003,
	ModelList:     20004,
}
