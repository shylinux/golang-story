package enums

var Errors = struct {
	Success int64
	Unknown int64

	AuthFailure       int64
	NotFoundProxy     int64
	InvalidParams     int64
	NotFoundUser      int64
	IncorrectPassword int64
	AlreadyExists     int64

	ModelCreate int64
	ModelRemove int64
	ModelModify int64
	ModelSearch int64
	ModelInfo   int64
	ModelList   int64
}{
	Success: 0,
	Unknown: 10000,

	AuthFailure:       10001,
	NotFoundProxy:     10002,
	InvalidParams:     10003,
	NotFoundUser:      10004,
	IncorrectPassword: 10005,
	AlreadyExists:     10006,

	ModelCreate: 20001,
	ModelRemove: 20002,
	ModelModify: 20003,
	ModelSearch: 20004,
	ModelInfo:   20005,
	ModelList:   20006,
}
