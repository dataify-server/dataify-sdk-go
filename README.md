# Dataify Go SDK

Generated Go SDK for Dataify tools.

This SDK exposes Go methods for Dataify API Key based runtime tools. SERP, Web Unlocker, and Scraper Builder tools call public Dataify HTTP APIs.

## Install

```bash
go get github.com/dataify-server/dataify-sdk-go
```

## Usage

```go
token := dataify.TokenFromEnv()
client := dataify.NewClient(token)
resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{Q: "coffee", JSON: "1"})
```

## Scraper Task Status

```go
status, err := client.Scraper.TaskStatus(ctx, "task_id_here")
if err != nil {
    // HTTP 400: task does not exist or is not owned by this API key.
    // HTTP 403: missing task_id or invalid API key.
}
fmt.Println(status.Data.Status) // 处理中, 成功, or 失败
```

## Scraper Task Download

```go
result, err := client.Scraper.DownloadTaskResult(ctx, "task_id_here")
// result is a RawResponse containing the JSON result.

response, err := client.Scraper.DownloadTaskFile(ctx, "task_id_here", dataify.ScraperDownloadXLSX)
if err == nil {
    defer response.Body.Close()
    // Stream or save the CSV/XLSX body.
}
```

## AI Web

AI Web provides synchronous search, map, and extract calls. Search supports web
(`AIWebUniqueIDWeb`), shopping, news, and scholar modes. The SDK defaults
to `https://aiwebapi.dataify.com`; set `AI_WEB_API_BASE_URL` or pass
`dataify.WithAIWebBaseURL` only when you need a test or custom gateway.

```go
client := dataify.NewClient(token)
result, err := client.AIWeb.Search(ctx, dataify.AIWebSearchRequest{
    UniqueID: dataify.AIWebUniqueIDWeb,
    Query:    "OpenAI",
    Source:   dataify.AIWebSourceWeb,
})
// result.RequestID comes from the X-Request-Id response header.
fmt.Println(result.String())
```

See [docs/AI_WEB.md](docs/AI_WEB.md) for the full parameter and download
reference. AI Web methods are separate from the generated runtime tools and are
not included in the tool count.

Set `DATAIFY_API_TOKEN` for new integrations. `DATAIFY_TOKEN` and
`DATAIFY_API_KEY` remain supported as compatibility aliases.
When a legacy variable supplies the credential, the SDK writes a one-time
migration warning to standard error without including the credential value.

Detailed Chinese usage guide: [docs/USAGE.zh-CN.md](docs/USAGE.zh-CN.md).

## Tool Count

72 API Key runtime tools are documented in docs/tools.md and docs/tools.yaml.
Task status is a separate Scraper runtime method and is not included in that count.
Task downloads are also separate Scraper runtime methods and are not included in that count.
