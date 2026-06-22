# Dataify Go SDK

Generated Go SDK for Dataify tools.

This SDK exposes Go methods for Dataify API Key based runtime tools. SERP, Web Unlocker, and Scraper Builder tools call public Dataify HTTP APIs.

## Install

```bash
go get github.com/dataify-server/dataify-sdk-go
```

## Usage

```go
client := dataify.NewClient("YOUR_DATAIFY_API_TOKEN")
resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{Q: "coffee", JSON: "1"})
```

## Tool Count

71 API Key runtime tools are documented in docs/tools.md and docs/tools.yaml.
