package http

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gogolook/domain"
	"gogolook/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HttpSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	e               *echo.Echo
	handler         ServerInterface
	mockTaskService *service.MockTaskServiceInterface
}

// TestSuiteInitTask is only for development and useCase, remove it in production
func TestSuiteInitHttp(t *testing.T) {
	suite.Run(t, new(HttpSuite))
}
func (t *HttpSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.e = echo.New()
	t.mockTaskService = service.NewMockTaskServiceInterface(t.mockCtrl)
	t.handler = NewRestfulServer(t.mockTaskService)
	RegisterHandlers(t.e.Group(""), t.handler)
}

func (t *HttpSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *HttpSuite) Test_get_all_task() {
	t.mockTaskService.EXPECT().FindAllTasks(gomock.Any()).
		Return([]domain.Task{}, nil).Times(1)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(req, rec)

	t.NoError(t.handler.GetTasks(c))
	t.Equal(http.StatusOK, rec.Code)
}

func (t *HttpSuite) Test_create_task() {
	testCase := []struct {
		input    string
		isError  bool
		httpCode int
	}{
		{`{"status":0}`, true, http.StatusBadRequest},
		{`{"name":"test","status":3}`, true, http.StatusBadRequest},
		{`{"name":"test","status":0}`, false, http.StatusCreated},
	}

	for _, currTestCase := range testCase {
		if currTestCase.isError {
			t.mockTaskService.EXPECT().CreateTask(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(domain.Task{}, nil).Times(0)
		} else {
			t.mockTaskService.EXPECT().CreateTask(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(domain.Task{}, nil).Times(1)
		}

		mockBody := []byte(currTestCase.input)
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(mockBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := t.e.NewContext(req, rec)
		err := t.handler.PostTask(c)
		if currTestCase.isError {
			var echoErr *echo.HTTPError
			t.True(errors.As(err, &echoErr))
			echoErr = err.(*echo.HTTPError)
			t.Error(err)
			t.Equal(currTestCase.httpCode, echoErr.Code)
		} else {
			t.NoError(err)
			t.Equal(currTestCase.httpCode, rec.Code)
		}
	}
}

func (t *HttpSuite) Test_update_task() {

	testCase := []struct {
		uuid       string
		updateData string
		isError    bool
		httpCode   int
	}{
		{"7e9d5dc8-d7ec-4bd2-8333-0bfbaef0a37d", `{"status":0}`, false, http.StatusOK},
		{"7e9d5dc8-d7ec-4bd2-8333-0bfbaef0a37d", `{"status":3}`, true, http.StatusBadRequest},
	}

	for _, currTestCase := range testCase {
		if currTestCase.isError {
			t.mockTaskService.EXPECT().UpdateTaskByID(gomock.Any(), currTestCase.uuid, gomock.Any()).
				Return(domain.Task{}, nil).Times(0)
		} else {
			t.mockTaskService.EXPECT().UpdateTaskByID(gomock.Any(), currTestCase.uuid, gomock.Any()).
				Return(domain.Task{}, nil).Times(1)
		}

		mockBody := []byte(currTestCase.updateData)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/tasks/%s", currTestCase.uuid), bytes.NewReader(mockBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := t.e.NewContext(req, rec)

		parseUUID := uuid.MustParse(currTestCase.uuid)
		t.NoError(t.handler.PutTask(c, parseUUID))
		t.Equal(http.StatusOK, rec.Code)
	}
}

func (t *HttpSuite) Test_delete_task() {
	uuidString := "7e9d5dc8-d7ec-4bd2-8333-0bfbaef0a37d"
	t.mockTaskService.EXPECT().DeleteTaskByID(gomock.Any(), uuidString).
		Return(nil).Times(1)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/tasks/%s", uuidString), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(req, rec)

	parseUUID := uuid.MustParse(uuidString)
	t.NoError(t.handler.DeleteTask(c, parseUUID))
	t.Equal(http.StatusNoContent, rec.Code)
}
