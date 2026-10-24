package dataify

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type chatGPTTransport struct {
	request *http.Request
	body    string
}

func (t *chatGPTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	if req.Body != nil {
		body, _ := io.ReadAll(req.Body)
		t.body = string(body)
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{"code":200,"data":{}}`)),
		Request:    req,
	}, nil
}

func TestChatGPTAnswerByURLBuildsBuilderPayload(t *testing.T) {
	transport := &chatGPTTransport{}
	client := NewClient("api key? &", WithHTTPClient(&http.Client{Transport: transport}))

	_, err := client.ChatGPT.Answer(context.Background(), ChatGPTAnswerRequest{
		SpiderID:   "chatgpt_answer_by-url",
		ChatGPTURL: "https://chatgpt.com/?q=pizza",
	})
	if err != nil {
		t.Fatalf("Answer() error = %v", err)
	}
	if transport.request.Method != http.MethodPost {
		t.Fatalf("method = %s, want %s", transport.request.Method, http.MethodPost)
	}
	if transport.request.URL.Path != "/builder" {
		t.Fatalf("path = %q, want /builder", transport.request.URL.Path)
	}
	if got := transport.request.URL.Query().Get("platform"); got != "1" {
		t.Fatalf("platform = %q, want 1", got)
	}
	if got := transport.request.Header.Get("Authorization"); got != "Bearer api key? &" {
		t.Fatalf("Authorization = %q", got)
	}
	form, err := url.ParseQuery(transport.body)
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got := form.Get("spider_name"); got != "chatgpt.com" {
		t.Fatalf("spider_name = %q, want chatgpt.com", got)
	}
	if got := form.Get("spider_id"); got != "chatgpt_answer_by-url" {
		t.Fatalf("spider_id = %q, want chatgpt_answer_by-url", got)
	}
	if got := form.Get("spider_errors"); got != "true" {
		t.Fatalf("spider_errors = %q, want true", got)
	}
	if got := form.Get("file_name"); got != "{{TasksID}}" {
		t.Fatalf("file_name = %q, want {{TasksID}}", got)
	}
	params := form.Get("spider_parameters")
	if !strings.Contains(params, `"chatgpt_url":"https://chatgpt.com/?q=pizza"`) {
		t.Fatalf("spider_parameters = %q, want chatgpt_url", params)
	}
	if strings.Contains(params, "search_terms") {
		t.Fatalf("spider_parameters should omit search_terms for by-url: %q", params)
	}
}

func TestChatGPTAnswerDefaultsToKeywordsSpider(t *testing.T) {
	transport := &chatGPTTransport{}
	client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

	_, err := client.ChatGPT.Answer(context.Background(), ChatGPTAnswerRequest{})
	if err != nil {
		t.Fatalf("Answer() error = %v", err)
	}
	form, err := url.ParseQuery(transport.body)
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got := form.Get("spider_id"); got != "chatgpt_answer_by-keywords" {
		t.Fatalf("spider_id = %q, want chatgpt_answer_by-keywords", got)
	}
}

func TestChatGPTAnswerRequiresLocalToken(t *testing.T) {
	t.Setenv(testStandardTokenEnv, "")
	t.Setenv(testLegacyTokenEnv, "")
	t.Setenv(testLegacyAPIKeyEnv, "")
	client := NewClient("")

	_, err := client.ChatGPT.Answer(context.Background(), ChatGPTAnswerRequest{})
	if err == nil {
		t.Fatal("Answer() error = nil, want ErrMissingToken")
	}
}
