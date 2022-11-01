// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	runner "github.com/caraml-dev/turing/engines/experiment/runner"
)

// ExperimentRunner is an autogenerated mock type for the ExperimentRunner type
type ExperimentRunner struct {
	mock.Mock
}

// GetTreatmentForRequest provides a mock function with given fields: header, payload, options
func (_m *ExperimentRunner) GetTreatmentForRequest(header http.Header, payload []byte, options runner.GetTreatmentOptions) (*runner.Treatment, error) {
	ret := _m.Called(header, payload, options)

	var r0 *runner.Treatment
	if rf, ok := ret.Get(0).(func(http.Header, []byte, runner.GetTreatmentOptions) *runner.Treatment); ok {
		r0 = rf(header, payload, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runner.Treatment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.Header, []byte, runner.GetTreatmentOptions) error); ok {
		r1 = rf(header, payload, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
