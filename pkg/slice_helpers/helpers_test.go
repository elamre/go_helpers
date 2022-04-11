package slice_helpers

import "testing"

func TestListEquals(t *testing.T) {
	if !ListEquals[int](nil, nil) {
		t.Error("nil lists should be equal")
	}
	if ListEquals[int](nil, []int{1, 2, 3}) {
		t.Error("one is nil should not equal")
	}
	if ListEquals[int]([]int{1, 2, 3}, nil) {
		t.Error("one is nil should not equal")
	}
	if !ListEquals[int]([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Error("Both should be equal")
	}
	if ListEquals[int]([]int{2, 2, 3}, []int{1, 2, 3}) {
		t.Error("One is not equal")
	}
	if ListEquals[int]([]int{1, 2, 3, 4}, []int{1, 2, 3}) {
		t.Error("Different lengths should not be equal")
	}
	if ListEquals[int]([]int{1, 2, 3}, []int{1, 2, 3, 4}) {
		t.Error("Different lengths should not be equal")
	}
}

func TestRemoveFromList(t *testing.T) {
	var RemoveResult []int
	RemoveResult = RemoveFromList[int](1, []int{1, 2, 3, 4})
	if !ListEquals[int](RemoveResult, []int{2, 3, 4}) {
		t.Errorf("Failed to remove first element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromList[int](2, []int{1, 2, 3, 4})
	if !ListEquals[int](RemoveResult, []int{1, 3, 4}) {
		t.Errorf("Failed to remove second element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromList[int](3, []int{1, 2, 3, 4})
	if !ListEquals[int](RemoveResult, []int{1, 2, 4}) {
		t.Errorf("Failed to remove third element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromList[int](4, []int{1, 2, 3, 4})
	if !ListEquals[int](RemoveResult, []int{1, 2, 3}) {
		t.Errorf("Failed to remove fourth element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromList[int](4, []int{4})
	if !ListEquals[int](RemoveResult, []int{}) {
		t.Errorf("Failed to remove the only element: %v", RemoveResult)
	}

}

func TestRemoveFromListEquals(t *testing.T) {
	var RemoveResult []int
	RemoveResult = RemoveFromListEquals[int]([]int{1, 2, 3, 4}, func(i int) bool { return i == 1 })
	if !ListEquals[int](RemoveResult, []int{2, 3, 4}) {
		t.Errorf("Failed to remove first element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromListEquals[int]([]int{1, 2, 3, 4}, func(i int) bool { return i == 2 })
	if !ListEquals[int](RemoveResult, []int{1, 3, 4}) {
		t.Errorf("Failed to remove second element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromListEquals[int]([]int{1, 2, 3, 4}, func(i int) bool { return i == 3 })
	if !ListEquals[int](RemoveResult, []int{1, 2, 4}) {
		t.Errorf("Failed to remove third element: %v", RemoveResult)
	}

	RemoveResult = RemoveFromListEquals[int]([]int{1, 2, 3, 4}, func(i int) bool { return i == 4 })
	if !ListEquals[int](RemoveResult, []int{1, 2, 3}) {
		t.Errorf("Failed to remove fourth element: %v", RemoveResult)
	}

	if RemoveFromListEquals[int]([]int{1, 2, 3, 4}, func(i int) bool { return i == 5 }) != nil {
		t.Error("We should have nil here, as the result is not in the array")
	}

	RemoveResult = RemoveFromListEquals[int]([]int{1}, func(i int) bool { return i == 1 })
	if !ListEquals[int](RemoveResult, []int{}) {
		t.Errorf("Failed to remove the only element: %v", RemoveResult)
	}
}
