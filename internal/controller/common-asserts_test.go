package controller

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

type assertSuccessCase struct {
	testName      string
	res           *http.Response
	req           *http.Request
	expectedError error
}

func TestAssertSuccessCase(t *testing.T) {
	testCases := []assertSuccessCase{
		{
			testName: "Successfully AssertSuccess",
			res:      &http.Response{StatusCode: http.StatusOK},
			req:      &http.Request{},
		},
		{
			testName:      "AssertSuccess throws error when it receives a different status than 200",
			res:           &http.Response{StatusCode: http.StatusBadRequest},
			req:           &http.Request{},
			expectedError: errors.New("error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			err := AssertSuccess(test.res, test.req)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

type assertInternalServerErrorCase struct {
	testName      string
	res           *http.Response
	req           *http.Request
	expectedError error
}

func TestAssertInternalServerErrorCase(t *testing.T) {
	testCases := []assertInternalServerErrorCase{
		{
			testName: "Successfully InternalServerError",
			res:      &http.Response{StatusCode: http.StatusInternalServerError},
			req:      &http.Request{},
		},
		{
			testName:      "InternalServerError throws error when it receives a different status than 500",
			res:           &http.Response{StatusCode: http.StatusOK},
			req:           &http.Request{},
			expectedError: errors.New("error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			err := AssertInternalServerError(test.res, test.req)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

type assertBadRequestCase struct {
	testName      string
	res           *http.Response
	req           *http.Request
	expectedError error
}

func TestAssertBadRequestCase(t *testing.T) {
	testCases := []assertBadRequestCase{
		{
			testName: "Successfully BadRequest",
			res:      &http.Response{StatusCode: http.StatusBadRequest},
			req:      &http.Request{},
		},
		{
			testName:      "BadRequest throws error when it receives a different status than 500",
			res:           &http.Response{StatusCode: http.StatusOK},
			req:           &http.Request{},
			expectedError: errors.New("error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			err := AssertBadRequest(test.res, test.req)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

type assertNotFoundCase struct {
	testName      string
	res           *http.Response
	req           *http.Request
	expectedError error
}

func TestAssertNotFoundCase(t *testing.T) {
	testCases := []assertNotFoundCase{
		{
			testName: "Successfully NotFound",
			res:      &http.Response{StatusCode: http.StatusNotFound},
			req:      &http.Request{},
		},
		{
			testName:      "NotFound throws error when it receives a different status than 500",
			res:           &http.Response{StatusCode: http.StatusOK},
			req:           &http.Request{},
			expectedError: errors.New("error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			err := AssertNotFound(test.res, test.req)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
