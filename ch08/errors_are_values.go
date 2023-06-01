package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

// StatusErr : Custom defined error struct
type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

// Using the struct to find more details about the error

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}

	return data, nil
}

func getData(file string) ([]byte, error) {
	// Some business logic
	fmt.Println("File name: " + file)
	return []byte{}, nil
}

func login(uid, pwd string) error {
	// Some business logic
	fmt.Println("uid " + uid)
	return nil
}
