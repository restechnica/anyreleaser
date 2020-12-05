package mocks

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/stretchr/testify/mock"
)

type MockCommander struct {
	mock.Mock
}

func NewMockCommander() *MockCommander {
	return &MockCommander{}
}

func (mock *MockCommander) Output(name string, arg ...string) (string, error) {
	args := mock.Called(name, arg)
	return args.String(0), args.Error(1)
}

func (mock *MockCommander) Run(name string, arg ...string) error {
	args := mock.Called(name, arg)
	return args.Error(0)
}

type MockConfigLoader struct {
	mock.Mock
}

func NewMockConfigLoader() *MockConfigLoader {
	return &MockConfigLoader{}
}

func (mock *MockConfigLoader) Load(path string) (cfg config.Root, err error) {
	var args = mock.Called(path)
	return args.Get(0).(config.Root), args.Error(1)
}

func (mock *MockConfigLoader) Overload(path string, cfg config.Root) (config.Root, error) {
	var args = mock.Called(path, cfg)
	return args.Get(0).(config.Root), args.Error(1)
}

type MockFlowPipe struct {
	mock.Mock
}

type MockGitService struct {
	mock.Mock
}

func NewMockGitService() *MockGitService {
	return &MockGitService{}
}

func (mock *MockGitService) CreateTag(tag string) (err error) {
	args := mock.Called(tag)
	return args.Error(0)
}

func (mock *MockGitService) GetLatestCommitMessage() (message string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

func (mock *MockGitService) GetTag() (output string) {
	args := mock.Called()
	return args.String(0)
}

func NewMockFlowPipe() *MockFlowPipe {
	return &MockFlowPipe{}
}

func (mock *MockFlowPipe) Run(ctx *app.Context) error {
	var args = mock.Called(ctx)
	return args.Error(0)
}

type MockSemverStrategy struct {
	mock.Mock
}

func NewMockSemverStrategy() *MockSemverStrategy {
	return &MockSemverStrategy{}
}

func (mock *MockSemverStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	args := mock.Called(targetVersion)
	return args.String(0), args.Error(1)
}
