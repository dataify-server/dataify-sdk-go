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

Set `DATAIFY_API_TOKEN` for new integrations. `DATAIFY_TOKEN` and
`DATAIFY_API_KEY` remain supported as compatibility aliases.
When a legacy variable supplies the credential, the SDK writes a one-time
migration warning to standard error without including the credential value.

Detailed Chinese usage guide: [docs/USAGE.zh-CN.md](docs/USAGE.zh-CN.md).

## Tool Count

71 API Key runtime tools are documented in docs/tools.md and docs/tools.yaml.
Task status is a separate Scraper runtime method and is not included in that count.
Task downloads are also separate Scraper runtime methods and are not included in that count.
