package main

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	// do something here
	return nil
}

func main() {
	err := MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	if err != nil {
		return
	}
	err = MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
	if err != nil {
		return
	}
}
