package controller

import (
	"errors"
	"net/http"
)

// AssertSuccess - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusOK
func AssertSuccess(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusOK {
		return errors.New("error")
	}

	return nil
}

// AssertInternalServerError - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusInternalServerError
func AssertInternalServerError(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusInternalServerError {
		return errors.New("error")
	}

	return nil
}

// AssertBadRequest - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusBadRequest
func AssertBadRequest(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusBadRequest {
		return errors.New("error")
	}

	return nil
}

// AssertNotFound - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusNotFound
func AssertNotFound(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusNotFound {
		return errors.New("error")
	}

	return nil
}
