package main

import "fmt"

type CustomErr struct {
	Status  Status
	Message string
	err     error
}

func (ce CustomErr) Error() string {
	return ce.Message
}

/*
If you want to wrap an error with your `custom error type`,
your error type needs to implement the method Unwrap.
*/

func (ce CustomErr) Unwrap() error {
	return ce.err
}

func LoginAndGetDataOfUser(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, CustomErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			err:     err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, CustomErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			err:     err,
		}
	}
	return data, nil
}
