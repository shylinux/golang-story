package check

func Assert(err error) {
	if err == nil {
		return
	}
	panic(err)
}
