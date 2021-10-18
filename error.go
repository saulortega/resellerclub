package resellerclub

import (
	"errors"
	"strings"
)

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrNoTLDsSelected     = errors.New("No TLDs are selected")
)

type Error struct {
	err string
	res interface{}
}

func (e Error) Error() string {
	return e.err
}

func (e Error) Response() interface{} {
	return e.res
}

func (e Error) Is(target error) bool {
	return e.err == target.Error()
}

func checkResponseError(mapResp map[string]interface{}) error {
	status, ok := mapResp["status"].(string)
	if ok {
		status = strings.ToLower(status)
		msg, ok := mapResp["message"].(string)
		if ok && len(msg) > 0 && status == "error" {
			return Error{msg, mapResp}
		}

		return somethingWentWrong(mapResp)
	}

	errorvalue, ok := mapResp["errorvalue"].(map[string]interface{})
	if ok {
		msg, ok := errorvalue["error"].(string)
		if ok && len(msg) > 0 {
			return Error{msg, mapResp}
		}

		return somethingWentWrong(mapResp)
	}

	return nil
}

func somethingWentWrong(res interface{}) error {
	return Error{
		ErrSomethingWentWrong.Error(),
		res,
	}
}
