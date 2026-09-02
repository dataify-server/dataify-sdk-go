package dataify

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ScraperTaskStatus string

type ScraperDownloadType string

const (
	ScraperTaskProcessing ScraperTaskStatus = "处理中"
	ScraperTaskSuccess    ScraperTaskStatus = "成功"
	ScraperTaskFailed     ScraperTaskStatus = "失败"
)

const (
	ScraperDownloadJSON ScraperDownloadType = "json"
	ScraperDownloadCSV  ScraperDownloadType = "csv"
	ScraperDownloadXLSX ScraperDownloadType = "xlsx"
)

var ErrUnsupportedDownloadType = errors.New("dataify: unsupported scraper download type")

type TaskStatusResponse struct {
	Data TaskStatusData `json:"data"`
	Code int            `json:"code"`
}

type TaskStatusData struct {
	TaskID string            `json:"task_id"`
	Status ScraperTaskStatus `json:"status"`
}

type ScraperService struct {
	client *Client
}

// TaskStatus queries a Scraper Builder task once. The service returns HTTP 400
// for an unknown or unauthorized task and HTTP 403 for missing parameters or an
// invalid API key, both as APIError values.
func (s *ScraperService) TaskStatus(ctx context.Context, taskID string) (*TaskStatusResponse, error) {
	if s.client.token == "" {
		return nil, ErrMissingToken
	}

	query := url.Values{}
	query.Set("api_key", s.client.token)
	query.Set("task_id", strings.TrimSpace(taskID))
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		strings.TrimRight(s.client.scraperBaseURL, "/")+"/task_status?"+query.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	raw, err := s.client.do(req)
	if err != nil {
		return nil, err
	}
	var response TaskStatusResponse
	if err := raw.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DownloadTaskResult downloads a task's JSON result and normalizes it as a RawResponse.
func (s *ScraperService) DownloadTaskResult(ctx context.Context, taskID string) (*RawResponse, error) {
	response, err := s.DownloadTaskFile(ctx, taskID, ScraperDownloadJSON)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return NewRawResponse(body)
}

// DownloadTaskFile downloads a task result in JSON, CSV, or XLSX format. The
// caller owns and must close the returned response body.
func (s *ScraperService) DownloadTaskFile(ctx context.Context, taskID string, format ScraperDownloadType) (*http.Response, error) {
	if s.client.token == "" {
		return nil, ErrMissingToken
	}
	if !isSupportedDownloadType(format) {
		return nil, ErrUnsupportedDownloadType
	}

	query := url.Values{}
	query.Set("api_key", s.client.token)
	query.Set("task_id", strings.TrimSpace(taskID))
	query.Set("type", string(format))
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		strings.TrimRight(s.client.scraperBaseURL, "/")+"/download?"+query.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		return response, nil
	}

	body, readErr := io.ReadAll(response.Body)
	response.Body.Close()
	if readErr != nil {
		return nil, readErr
	}
	return nil, &APIError{StatusCode: response.StatusCode, Body: body}
}

func isSupportedDownloadType(format ScraperDownloadType) bool {
	switch format {
	case ScraperDownloadJSON, ScraperDownloadCSV, ScraperDownloadXLSX:
		return true
	default:
		return false
	}
}
