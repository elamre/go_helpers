package misc

import "fmt"

type testStruct struct {
	number int
}

func (t testStruct) String() string {
	return fmt.Sprintf("%d", t.number)
}

func (t testStruct) IsError() error {
	if t.number == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func (t testStruct) IsErrorAndData() (int, error) {
	if t.number == 0 {
		return 0, fmt.Errorf("error")
	}
	return t.number, nil
}

func ExampleCast() {
	t := &testStruct{number: 4}
	fmt.Println(Cast[*testStruct](t).String())
	// Output: 4
}

func ExampleCheckError() {
	t := &testStruct{number: 4}
	CheckError(t.IsError())
	// Output:
}

func ExampleCheckErrorRetVal() {
	t := &testStruct{number: 4}
	number := CheckErrorRetVal[int](t.IsErrorAndData())
	fmt.Printf("%d", number)
	// Output: 4
}
