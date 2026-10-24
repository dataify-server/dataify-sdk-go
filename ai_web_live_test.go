package dataify

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

// TestAIWebLiveSmoke is an opt-in integration test against a real AI Web
// gateway. It is skipped unless DATAIFY_AI_WEB_LIVE=1. Set
// AI_WEB_API_BASE_URL and DATAIFY_API_TOKEN before enabling it.
func TestAIWebLiveSmoke(t *testing.T) {
	if strings.TrimSpace(os.Getenv("DATAIFY_AI_WEB_LIVE")) != "1" {
		t.Skip("set DATAIFY_AI_WEB_LIVE=1 to run the live AI Web smoke test")
	}

	baseURL := strings.TrimSpace(os.Getenv("AI_WEB_API_BASE_URL"))
	if baseURL == "" {
		t.Fatal("AI_WEB_API_BASE_URL is required for the live AI Web smoke test")
	}
	token := strings.TrimSpace(os.Getenv("DATAIFY_API_TOKEN"))
	if token == "" {
		t.Fatal("DATAIFY_API_TOKEN is required for the live AI Web smoke test")
	}

	client := NewClient(token, WithAIWebBaseURL(baseURL), WithTimeout(90*time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	search, err := client.AIWeb.Search(ctx, AIWebSearchRequest{
		UniqueID:   AIWebUniqueIDNews,
		Query:      "OpenAI",
		MaxResults: IntPtr(2),
		Source:     AIWebSourceNews,
	})
	if err != nil {
		t.Fatalf("live search failed: %v", err)
	}
	searchItems, err := search.OutputsList()
	if err != nil {
		t.Fatalf("live search outputs are not a JSON array: %v (%s)", err, search.String())
	}
	t.Logf("search ok: request_id=%q outputs=%d", search.RequestID, len(searchItems))

	mapResult, err := client.AIWeb.Map(ctx, AIWebMapRequest{
		UniqueID: AIWebUniqueIDMap,
		URL:      "https://example.com",
		Limit:    IntPtr(5),
	})
	if err != nil {
		t.Fatalf("live map failed: %v", err)
	}
	mapItems, err := mapResult.OutputsList()
	if err != nil {
		t.Fatalf("live map outputs are not a JSON array: %v (%s)", err, mapResult.String())
	}
	t.Logf("map ok: request_id=%q outputs=%d", mapResult.RequestID, len(mapItems))

	extract, err := client.AIWeb.Extract(ctx, AIWebExtractRequest{
		UniqueID: AIWebUniqueIDExtract,
		URL:      "https://example.com",
		Formats:  []string{AIWebFormatMarkdown},
	})
	if err != nil {
		t.Fatalf("live extract failed: %v", err)
	}
	extractItems, err := extract.OutputsList()
	if err != nil {
		t.Fatalf("live extract outputs are not a JSON array: %v (%s)", err, extract.String())
	}
	t.Logf("extract ok: request_id=%q outputs=%d", extract.RequestID, len(extractItems))

	if extract.RequestID != "" {
		downloaded, err := client.AIWeb.DownloadResult(ctx, extract.RequestID)
		if err != nil {
			t.Logf("download attempt for request_id=%q failed: %v", extract.RequestID, err)
			return
		}
		t.Logf("download ok: bytes=%d", len(downloaded.Bytes()))
	}
}
