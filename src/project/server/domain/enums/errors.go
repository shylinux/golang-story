package enums

var Errors = struct {
	Success     int
	Unknown     int
	ModelCreate int
	ModelRemove int
	ModelInfo   int
	ModelList   int
}{
	Success:     0,
	Unknown:     10000,
	ModelCreate: 10001,
	ModelRemove: 10002,
	ModelInfo:   10003,
	ModelList:   10004,
}
