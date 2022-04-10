package slice_helpers

func ListEquals[t comparable](listOne []t, listTwo []t) bool {
	if len(listOne) != len(listTwo) {
		return false
	}
	if listOne == nil && listTwo == nil {
		return true
	}
	if len(listOne) == 0 {
		return false
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
		return make([]t, 0)
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
