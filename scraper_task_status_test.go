package dataify

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type taskStatusTransport struct {
	request    *http.Request
	statusCode int
	body       string
}

func (t *taskStatusTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	return &http.Response{
		StatusCode: t.statusCode,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(t.body)),
		Request:    req,
	}, nil
}

func TestScraperTaskStatusSendsDocumentedQueryAndDecodesChineseStatus(t *testing.T) {
	transport := &taskStatusTransport{
		statusCode: http.StatusOK,
		body:       `{"data":{"task_id":"task id? &","status":"处理中"},"code":200}`,
	}
	client := NewClient("api key? &", WithHTTPClient(&http.Client{Transport: transport}))

	response, err := client.Scraper.TaskStatus(context.Background(), " task id? & ")
	if err != nil {
		t.Fatalf("TaskStatus() error = %v", err)
	}
	if response.Code != http.StatusOK {
		t.Fatalf("response code = %d, want %d", response.Code, http.StatusOK)
	}
	if response.Data.TaskID != "task id? &" {
		t.Fatalf("task ID = %q, want %q", response.Data.TaskID, "task id? &")
	}
	if response.Data.Status != ScraperTaskProcessing {
		t.Fatalf("status = %q, want %q", response.Data.Status, ScraperTaskProcessing)
	}
	if transport.request.Method != http.MethodGet {
		t.Fatalf("method = %s, want %s", transport.request.Method, http.MethodGet)
	}
	if transport.request.URL.Path != "/task_status" {
		t.Fatalf("path = %q, want /task_status", transport.request.URL.Path)
	}
	if got := transport.request.URL.Query().Get("api_key"); got != "api key? &" {
		t.Fatalf("api_key = %q, want %q", got, "api key? &")
	}
	if got := transport.request.URL.Query().Get("task_id"); got != "task id? &" {
		t.Fatalf("task_id = %q, want %q", got, "task id? &")
	}
	if got := transport.request.Header.Get("Authorization"); got != "" {
		t.Fatalf("Authorization = %q, want no header", got)
	}
}

func TestScraperTaskStatusDecodesAllDocumentedStatuses(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want ScraperTaskStatus
	}{
		{name: "processing", raw: "处理中", want: ScraperTaskProcessing},
		{name: "successful", raw: "成功", want: ScraperTaskSuccess},
		{name: "failed", raw: "失败", want: ScraperTaskFailed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &taskStatusTransport{
				statusCode: http.StatusOK,
				body:       `{"data":{"task_id":"task-123","status":"` + tt.raw + `"},"code":200}`,
			}
			client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

			response, err := client.Scraper.TaskStatus(context.Background(), "task-123")
			if err != nil {
				t.Fatalf("TaskStatus() error = %v", err)
			}
			if response.Data.Status != tt.want {
				t.Fatalf("status = %q, want %q", response.Data.Status, tt.want)
			}
		})
	}
}

func TestScraperTaskStatusPropagatesServiceHTTPFailures(t *testing.T) {
	tests := []struct {
		name     string
		taskID   string
		status   int
		body     string
		wantTask string
	}{
		{
			name:     "missing parameter returns forbidden",
			taskID:   "   ",
			status:   http.StatusForbidden,
			body:     `{"data":"missing task_id","code":403}`,
			wantTask: "",
		},
		{
			name:     "unknown or unauthorized task returns bad request",
			taskID:   "unknown-task",
			status:   http.StatusBadRequest,
			body:     `{"data":"Task_id is error!","code":400}`,
			wantTask: "unknown-task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &taskStatusTransport{statusCode: tt.status, body: tt.body}
			client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

			_, err := client.Scraper.TaskStatus(context.Background(), tt.taskID)
			var apiErr *APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("TaskStatus() error = %v, want APIError", err)
			}
			if apiErr.StatusCode != tt.status {
				t.Fatalf("status = %d, want %d", apiErr.StatusCode, tt.status)
			}
			if got := transport.request.URL.Query().Get("task_id"); got != tt.wantTask {
				t.Fatalf("task_id = %q, want %q", got, tt.wantTask)
			}
		})
	}
}

func TestScraperTaskStatusRejectsMissingLocalToken(t *testing.T) {
	t.Setenv(testStandardTokenEnv, "")
	t.Setenv(testLegacyTokenEnv, "")
	t.Setenv(testLegacyAPIKeyEnv, "")

	client := NewClient("")
	_, err := client.Scraper.TaskStatus(context.Background(), "task-123")
	if !errors.Is(err, ErrMissingToken) {
		t.Fatalf("TaskStatus() error = %v, want ErrMissingToken", err)
	}
}
