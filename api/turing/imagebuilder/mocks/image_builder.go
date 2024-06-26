// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	client "github.com/caraml-dev/mlp/api/client"
	imagebuilder "github.com/caraml-dev/turing/api/turing/imagebuilder"
	mock "github.com/stretchr/testify/mock"

	models "github.com/caraml-dev/turing/api/turing/models"
)

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: request
func (_m *ImageBuilder) BuildImage(request imagebuilder.BuildImageRequest) (string, error) {
	ret := _m.Called(request)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(imagebuilder.BuildImageRequest) (string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(imagebuilder.BuildImageRequest) string); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(imagebuilder.BuildImageRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImageBuildingJob provides a mock function with given fields: projectName, ensemblerName, ensemblerID, versionID
func (_m *ImageBuilder) DeleteImageBuildingJob(projectName string, ensemblerName string, ensemblerID models.ID, versionID string) error {
	ret := _m.Called(projectName, ensemblerName, ensemblerID, versionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, models.ID, string) error); ok {
		r0 = rf(projectName, ensemblerName, ensemblerID, versionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEnsemblerImage provides a mock function with given fields: project, ensembler
func (_m *ImageBuilder) GetEnsemblerImage(project *client.Project, ensembler *models.PyFuncEnsembler) (imagebuilder.EnsemblerImage, error) {
	ret := _m.Called(project, ensembler)

	var r0 imagebuilder.EnsemblerImage
	var r1 error
	if rf, ok := ret.Get(0).(func(*client.Project, *models.PyFuncEnsembler) (imagebuilder.EnsemblerImage, error)); ok {
		return rf(project, ensembler)
	}
	if rf, ok := ret.Get(0).(func(*client.Project, *models.PyFuncEnsembler) imagebuilder.EnsemblerImage); ok {
		r0 = rf(project, ensembler)
	} else {
		r0 = ret.Get(0).(imagebuilder.EnsemblerImage)
	}

	if rf, ok := ret.Get(1).(func(*client.Project, *models.PyFuncEnsembler) error); ok {
		r1 = rf(project, ensembler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImageBuildingJobStatus provides a mock function with given fields: projectName, ensemblerName, ensemblerID, versionID
func (_m *ImageBuilder) GetImageBuildingJobStatus(projectName string, ensemblerName string, ensemblerID models.ID, versionID string) imagebuilder.JobStatus {
	ret := _m.Called(projectName, ensemblerName, ensemblerID, versionID)

	var r0 imagebuilder.JobStatus
	if rf, ok := ret.Get(0).(func(string, string, models.ID, string) imagebuilder.JobStatus); ok {
		r0 = rf(projectName, ensemblerName, ensemblerID, versionID)
	} else {
		r0 = ret.Get(0).(imagebuilder.JobStatus)
	}

	return r0
}

type mockConstructorTestingTNewImageBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageBuilder creates a new instance of ImageBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageBuilder(t mockConstructorTestingTNewImageBuilder) *ImageBuilder {
	mock := &ImageBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
