package clientErrors

import "fmt"

type ClientSetupError struct{ Err error }

func (e ClientSetupError) Error() string {
	return fmt.Sprintf("client setup error: %s", e.Err)
}

type JoinPathError struct{ Err error }

func (e JoinPathError) Error() string {
	return fmt.Sprintf("join path error: %s", e.Err)
}

type RequestCreationError struct{ Err error }

func (e RequestCreationError) Error() string {
	return fmt.Sprintf("request creation error: %s", e.Err)
}

type RequestMarshalError struct{ Err error }

func (e RequestMarshalError) Error() string {
	return fmt.Sprintf("request marshal error: %s", e.Err)
}

type JsonAndBodyConflict struct{}

func (JsonAndBodyConflict) Error() string {
	return "json and body conflict"
}

type PrepareRequestBodyError struct{}

func (PrepareRequestBodyError) Error() string {
	return "prepare request body error"
}

type RunRequestError struct{ Err error }

func (e RunRequestError) Error() string {
	return fmt.Sprintf("prepare request error: %s", e.Err)
}

type SendRequestError struct{ Err error }

func (e SendRequestError) Error() string {
	return fmt.Sprintf("send request error: %s", e.Err)
}
