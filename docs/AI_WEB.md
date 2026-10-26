# AI Web API

The Go SDK exposes synchronous AI Web search, map, and extract calls through
`client.AIWeb`. All requests use the Dataify Bearer token. The SDK reads
the response header `X-Request-Id` and exposes it as `RequestID`.

## Basic usage

```go
client := dataify.NewClient(dataify.TokenFromEnv())

result, err := client.AIWeb.Search(ctx, dataify.AIWebSearchRequest{
    UniqueID: dataify.AIWebUniqueIDNews,
    Query:    "OpenAI",
    MaxResults: dataify.IntPtr(3),
    Source:   dataify.AIWebSourceNews,
})
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.RequestID)
fmt.Println(result.String())
```

## Methods

| SDK method | Upstream endpoint | Description |
|---|---|---|
| `client.AIWeb.Search(params)` | `POST /search` | Search web, news, shopping, or scholar results; optionally extract result content |
| `client.AIWeb.Map(params)` | `POST /map` | Discover page URLs, titles, and descriptions |
| `client.AIWeb.Extract(params)` | `POST /extract` | Extract markdown, raw HTML, clean HTML, or links |
| `client.AIWeb.Execute(operation, params)` | `POST /search`, `/map`, `/extract` | Call the selected operation and normalize outputs to an array |
| `client.AIWeb.ExecuteRaw(operation, params)` | `POST /search`, `/map`, `/extract` | Preserve the upstream outputs shape |
| `client.AIWeb.DownloadResult(requestID)` | `POST /download` | Download a retained AI Web result |

Search, Map, Extract, and Execute normalize `Outputs` to a JSON array: null
becomes an empty array, arrays remain arrays, and a single object is wrapped in
a one-element array. `ExecuteRaw` and `DownloadResult` preserve the upstream
shape.

`DownloadResult` returns the same `*AIWebResult` envelope as the other AI Web
methods. `RequestID` echoes the id you passed in, and `Outputs` contains the
retained raw JSON result.

## Validation

All AI Web methods validate the request before sending it. Invalid requests
return `*dataify.ValidationError` with `Parameter` set to the JSON field name.

- `unique_id`, `query`, and `url` are required where applicable.
- `max_results` must be 1..100; `limit` must be 1..100000.
- `time_range` accepts `hour`, `day`, `week`, `month`, or `year`. It is
  rejected for shopping and scholar search.
- `country` must be a two-letter code and is normalized to uppercase. It is
  rejected for scholar search.
- `domain_filter` accepts `all` or `domain`; `sitemap` accepts `skip`,
  `include`, or `only`.
- `formats` accepts `markdown`, `raw_html`, `clean_html`, or `links` without
  duplicates.
- `include_tags` and `exclude_tags` accept at most 50 items, each at most 2048
  bytes.
- `max_age_ms` must be 0..172800000; `timeout_ms` must be 1000..120000;
  `wait_for_ms` must be 0..30000 and at most half of `timeout_ms`.
- `extract_options` requires `extract_content=true`.
- `url` must be an absolute `http://` or `https://` URL without user info.
- `headers` must not contain `Host` or hop-by-hop headers.

For the built-in search tools, `source` is derived from `unique_id` when it is
omitted and must match when it is set.

## Common parameters

### Search

```go
result, err := client.AIWeb.Search(ctx, dataify.AIWebSearchRequest{
    UniqueID:       dataify.AIWebUniqueIDNews,
    Query:          "OpenAI",
    MaxResults:     dataify.IntPtr(10),
    Source:         dataify.AIWebSourceNews,
    Country:        "US",
    Language:       "en",
    TimeRange:      "week",
    ExtractContent: dataify.BoolPtr(true),
    ExtractOptions: &dataify.AIWebExtractOptions{
        Formats:         []string{dataify.AIWebFormatMarkdown},
        OnlyMainContent: dataify.BoolPtr(true),
        AdvancedMode:    dataify.BoolPtr(true),
    },
})
```

### Web search

`AIWebUniqueIDWeb` maps to `unique_id: search_web` and `source: web`. It
supports `time_range`, `country`, `language`, and the same extraction options
as the other search modes.

```go
result, err := client.AIWeb.Search(ctx, dataify.AIWebSearchRequest{
    UniqueID:   dataify.AIWebUniqueIDWeb,
    Query:      "OpenAI",
    MaxResults: dataify.IntPtr(10),
    Country:    "US",
    TimeRange:  "week",
})
```
The fixed search unique IDs are `AIWebUniqueIDWeb`,
`AIWebUniqueIDShopping`, `AIWebUniqueIDNews`, and `AIWebUniqueIDScholar`. A
custom gateway may still use a different `UniqueID`.

### Map

```go
result, err := client.AIWeb.Map(ctx, dataify.AIWebMapRequest{
    UniqueID:     dataify.AIWebUniqueIDMap,
    URL:          "https://docs.example.com",
    Country:      "US",
    DomainFilter: dataify.AIWebDomainFilterAll,
    Limit:        dataify.IntPtr(100),
    Sitemap:      dataify.AIWebSitemapInclude,
})
```

### Extract

```go
result, err := client.AIWeb.Extract(ctx, dataify.AIWebExtractRequest{
    UniqueID:        dataify.AIWebUniqueIDExtract,
    URL:             "https://example.com/article",
    Formats:         []string{dataify.AIWebFormatMarkdown, dataify.AIWebFormatLinks},
    OnlyMainContent: dataify.BoolPtr(true),
    Headers:         map[string]string{"User-Agent": "Dataify SDK"},
    TimeoutMS:       dataify.IntPtr(60000),
    AdvancedMode:    dataify.BoolPtr(true),
})
```

## Downloading results

Keep `RequestID` and use it while the result is retained:

```go
result, err := client.AIWeb.Extract(ctx, dataify.AIWebExtractRequest{
    UniqueID: dataify.AIWebUniqueIDExtract,
    URL:      "https://example.com/article",
    Formats:  []string{dataify.AIWebFormatMarkdown},
})
if err != nil {
    log.Fatal(err)
}
if result.RequestID != "" {
    downloaded, err := client.AIWeb.DownloadResult(ctx, result.RequestID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(downloaded.RequestID)
    fmt.Println(downloaded.String())
}
```

The downloaded payload is not re-normalized to an array. Use `ExecuteRaw` when
you need the same raw behavior for search, map, or extract.

## Errors

All SDK-originated errors implement `dataify.DataifyError`:

```go
var sdkErr dataify.DataifyError
if errors.As(err, &sdkErr) {
    // handle a Dataify SDK error
}
```

- HTTP status failures return `*dataify.APIError` with `StatusCode` and `Body`.
- Invalid local parameters return `*dataify.ValidationError` with `Parameter`
  and `Message`.
- Transport failures return `*dataify.NetworkError` with `Kind`. `errors.Is`
  matches `dataify.ErrNetwork`, `dataify.ErrTimeout`, or `dataify.ErrCanceled`.
- Missing API token returns `dataify.ErrMissingToken`.
- Disabling AI Web with an empty base URL returns `dataify.ErrMissingAIWebBaseURL`.
- Unsupported operations return `dataify.ErrUnsupportedAIWebOperation`.
- Undecodable successful responses return `dataify.ErrInvalidResponse`.

`NetworkError` preserves the underlying error, so
`errors.Is(err, context.DeadlineExceeded)` and `errors.As(err, &urlErr)` still
work.

## Gateway configuration

AI Web defaults to the production gateway `https://aiwebapi.dataify.com`.
For a test or custom gateway, set `AI_WEB_API_BASE_URL` before constructing
the client, or pass `dataify.WithAIWebBaseURL`:

```go
client := dataify.NewClient(
    dataify.TokenFromEnv(),
    dataify.WithAIWebBaseURL("https://aiwebapi.dataify.com"), // replace for test/custom
)
```

The environment variable takes precedence over the default; an explicit
`WithAIWebBaseURL` takes precedence over both. Passing an empty
`WithAIWebBaseURL("")` disables AI Web calls and makes them return
`ErrMissingAIWebBaseURL` before sending.

HTTP failures return `*dataify.APIError` with `StatusCode` and `Body`. A
missing API token returns `dataify.ErrMissingToken` before the request is sent.