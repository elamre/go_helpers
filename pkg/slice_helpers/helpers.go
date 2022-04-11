package slice_helpers

func ListEquals[t comparable](listOne []t, listTwo []t) bool {
	if len(listOne) != len(listTwo) {
		return false
	}
	if listOne == nil && listTwo == nil {
		return true
	}
	if len(listOne) == 0 {
		return true
	}
	for i := 0; i < len(listOne); i++ {
		if listOne[i] != listTwo[i] {
			return false
		}
	}
	return true
}

func RemoveFromList[t comparable](element t, list []t) []t {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 {
		if list[0] == element {
			return make([]t, 0)
		}
		return nil
	} else if list[0] == element {
		return list[1:]
	} else if list[len(list)-1] == element {
		return list[:len(list)-1]
	}
	for i := 1; i < len(list)-1; i++ {
		if list[i] == element {
			list = append(list[:i], list[i+1:]...)
			return list
		}
	}
	return nil
}

func RemoveFromListEquals[t any](list []t, equals func(t) bool) []t {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 {
		return make([]t, 0)
	}
	var i int
	var found bool
	for i = 0; i < len(list); i++ {
		if equals(list[i]) {
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	if len(list) == 1 {
		return make([]t, 0)
	} else if i == 0 {
		return list[1:]
	} else if i == len(list)-1 {
		return list[:len(list)-1]
	}
	list = append(list[:i], list[i+1:]...)
	return list
}
