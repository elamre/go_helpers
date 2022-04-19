package misc

// CheckError checks the return of a function returning an error. Panics if error != nil
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckErrorRetVal checks the error value of a function returning an error and data
// It's using generics for the return value type
func CheckErrorRetVal[t any](value t, err error) t {
	CheckError(err)
	return value
}

// Cast function casts an interface to a type using generics.
// returns obj.(t)
func Cast[t any](obj interface{}) t {
	return obj.(t)
}
