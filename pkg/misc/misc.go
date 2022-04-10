package misc

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrorRetVal[t any](value t, err error) t {
	CheckError(err)
	return value
}
