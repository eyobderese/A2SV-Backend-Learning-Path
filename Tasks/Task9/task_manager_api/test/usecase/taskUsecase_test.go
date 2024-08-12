package usecase_test

import (
	"testing"

	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/domain"
	"github.com/eyobderese/A2SV-Backend-Learning-Path/task_manager_api/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type taskUsecaseSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the funcionalities we need to test
	usecase domain.TaskUsecase

	// some helper function to clean-up any used tables
}

type taskRepositoryMock struct {
	mock.Mock
}

func (m *taskRepositoryMock) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *taskRepositoryMock) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) GetTaskById(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) UpdateTask(task domain.Task, id string) (domain.Task, error) {
	args := m.Called(task, id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *taskRepositoryMock) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (suite *taskUsecaseSuite) SetupSuite() {
	// this function runs once before all tests in the suite

	// some initialization setup

	repository := new(taskRepositoryMock)
	usecase := usecase.NewTaskUsecase(repository)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.usecase = usecase

}

func (suite *taskUsecaseSuite) TearDownSuite() {

}

func Test_taskUsecaseSuite(t *testing.T) {
	/// we still need this to run all tests in our suite
	suite.Run(t, &taskUsecaseSuite{})
}
