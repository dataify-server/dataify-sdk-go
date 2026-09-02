# Dataify Go SDK 使用文档

本文档说明如何在其他 Go 项目中安装、配置和调用 `github.com/dataify-server/dataify-sdk-go`。

注意：不要把真实 API Key 写进源码、README、Git 提交记录或日志。推荐通过环境变量读取。

## 1. 适用场景

这个 SDK 适合在 Go 项目里调用 Dataify 的抓取和搜索能力，包括：

- SERP 搜索引擎接口，例如 Google、Bing、Yandex、DuckDuckGo 等。
- Scraper Builder 抓取工具，例如 Amazon、YouTube、TikTok、LinkedIn、Walmart、Zillow 等平台。
- Web Unlocker 通用网页解锁接口。

当前 SDK 只保留 API Key 鉴权相关的抓取类工具，不包含用户信息、余额、统计、任务查询等需要 dashboard JWT 的接口。

## 2. 安装

在你的 Go 项目目录里执行：

```cmd
go get github.com/dataify-server/dataify-sdk-go@v0.1.0
```

如果你希望使用最新版本：

```cmd
go get github.com/dataify-server/dataify-sdk-go@latest
```

查看是否安装成功：

```cmd
go list -m all
```

输出里应该能看到：

```text
github.com/dataify-server/dataify-sdk-go v0.1.0
```

## 3. 环境变量配置

推荐把 API Token 放到环境变量 `DATAIFY_API_TOKEN`。这是所有新项目和文档中的标准变量名。

SDK 仍兼容旧变量 `DATAIFY_TOKEN` 和 `DATAIFY_API_KEY`。当没有显式传入可用 Token 时，优先级为：`DATAIFY_API_TOKEN`、`DATAIFY_TOKEN`、`DATAIFY_API_KEY`。不要在新项目中继续配置旧变量。

当 SDK 从旧变量读取，或传入值与旧变量当前值匹配时，会在标准错误中各提示一次迁移建议，例如：`dataify: DATAIFY_TOKEN 已兼容读取，建议迁移至 DATAIFY_API_TOKEN。`。提示不包含 Token 值；标准变量 `DATAIFY_API_TOKEN` 不会触发提示。

PowerShell：

```powershell
$env:DATAIFY_API_TOKEN="<api-key>"
```

CMD：

```cmd
set DATAIFY_API_TOKEN=<api-key>
```

Linux/macOS：

```bash
export DATAIFY_API_TOKEN="<api-key>"
```

这里的 `<api-key>` 是占位符，实际使用时替换成你自己的 Dataify API Key。不要把真实值提交到 Git。

## 4. 最小示例

新建 `main.go`：

```go
package main

import (
	"context"
	"fmt"
	"log"

	dataify "github.com/dataify-server/dataify-sdk-go"
)

func main() {
	token := dataify.TokenFromEnv()
	if token == "" {
		log.Fatal("DATAIFY_API_TOKEN is not set")
	}

	client := dataify.NewClient(token)
	ctx := context.Background()

	resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{
		Q:    "Dataify",
		JSON: "1",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())
}
```

运行：

```cmd
go run .
```

## 5. 客户端初始化

基础方式：

```go
client := dataify.NewClient(token)
```

设置超时时间：

```go
client := dataify.NewClient(token, dataify.WithTimeout(180*time.Second))
```

使用自定义 `http.Client`：

```go
httpClient := &http.Client{Timeout: 180 * time.Second}
client := dataify.NewClient(token, dataify.WithHTTPClient(httpClient))
```

如果你有代理、测试环境或私有网关，可以覆盖默认 Base URL：

```go
client := dataify.NewClient(
	token,
	dataify.WithSERPBaseURL("https://scraperapi.dataify.com"),
	dataify.WithScraperBaseURL("https://scraperapi.dataify.com"),
	dataify.WithWebUnlockerBaseURL("https://webunlocker.dataify.com"),
)
```

一般用户不需要修改 Base URL。

## 6. 返回值处理

所有接口返回：

```go
(*dataify.RawResponse, error)
```

`RawResponse` 常用方法：

```go
resp.String()  // 返回 JSON 字符串
resp.Bytes()   // 返回 JSON 字节
resp.Decode(&v) // 反序列化到结构体或 map
```

示例：

```go
var result map[string]any
if err := resp.Decode(&result); err != nil {
	log.Fatal(err)
}
fmt.Println(result)
```

## 7. SERP 示例

Google 搜索：

```go
resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{
	Q:    "Dataify",
	JSON: "1",
	GL:   "us",
	HL:   "en",
})
```

Bing 搜索：

```go
resp, err := client.Bing.Search(ctx, dataify.BingSearchRequest{
	Q:    "Dataify",
	JSON: "1",
})
```

Yandex 搜索：

```go
resp, err := client.Yandex.Search(ctx, dataify.YandexSearchRequest{
	Text: "Dataify",
	JSON: "1",
})
```

DuckDuckGo 搜索：

```go
resp, err := client.DuckDuckGo.Search(ctx, dataify.DuckDuckGoSearchRequest{
	Q:    "Dataify",
	JSON: "1",
})
```

更多 SERP 方法和参数见：

- [tools.md](tools.md)
- [tools.yaml](tools.yaml)

## 8. Scraper Builder 示例

Amazon 商品详情：

```go
resp, err := client.Amazon.Product(ctx, dataify.AmazonProductRequest{
	SpiderID: "amazon_product_by-url",
	URL:      "https://www.amazon.com/dp/B0BRXPR726",
	ZipCode: "94107",
})
```

YouTube 视频信息：

```go
resp, err := client.YouTube.Product(ctx, dataify.YouTubeProductRequest{
	VideoID: "8RePenzQH80",
})
```

Walmart 商品信息：

```go
resp, err := client.Walmart.Product(ctx, dataify.WalmartProductRequest{
	SpiderID: "walmart_product_by-sku",
	SKU:      "439179861",
})
```

Airbnb 房源搜索：

```go
resp, err := client.Airbnb.Product(ctx, dataify.AirbnbProductRequest{
	Searchurl: "https://www.airbnb.com/s/Greece/homes",
	Country:   "US",
})
```

Scraper Builder 类方法一般会提交采集任务，返回任务相关信息。不同工具的请求参数不同，完整参数以 [tools.md](tools.md) 和 [tools.yaml](tools.yaml) 为准。

## 9. Scraper 任务状态查询

Builder 任务返回 `task_id` 后，可以查询一次当前状态：

```go
status, err := client.Scraper.TaskStatus(ctx, "your_task_id")
if err != nil {
	var apiErr *dataify.APIError
	switch {
	case errors.As(err, &apiErr) && apiErr.StatusCode == 400:
		log.Fatal("task does not exist or is not owned by this API token")
	case errors.As(err, &apiErr) && apiErr.StatusCode == 403:
		log.Fatal("missing task_id or invalid API token")
	default:
		log.Fatal(err)
	}
}

fmt.Println(status.Data.TaskID)
fmt.Println(status.Data.Status) // 处理中、成功、失败
```

该方法请求 `GET https://scraperapi.dataify.com/task_status`，按接口约定以 `api_key` 和 `task_id` 查询参数鉴权和定位任务。它只查询一次，不会自动轮询或重试。任务不存在或不属于当前 API Token 时服务端返回 HTTP `400`；缺少参数或 API Token 无效时返回 HTTP `403`。

它是独立的 Scraper 运行方法，不属于 `tools.md` 与 `tools.yaml` 中的 71 个生成工具。

## 10. Scraper 任务结果下载

下载 JSON 结果时使用 `DownloadTaskResult`，返回已有的 `RawResponse`：

```go
result, err := client.Scraper.DownloadTaskResult(ctx, "your_task_id")
if err != nil {
	log.Fatal(err)
}
fmt.Println(result.String())
```

下载 CSV 或 XLSX 时使用 `DownloadTaskFile`，响应体由调用方关闭和保存：

```go
response, err := client.Scraper.DownloadTaskFile(ctx, "your_task_id", dataify.ScraperDownloadXLSX)
if err != nil {
	log.Fatal(err)
}
defer response.Body.Close()

file, err := os.Create("result.xlsx")
if err != nil {
	log.Fatal(err)
}
defer file.Close()
if _, err := io.Copy(file, response.Body); err != nil {
	log.Fatal(err)
}
```

支持 `dataify.ScraperDownloadJSON`、`dataify.ScraperDownloadCSV` 和 `dataify.ScraperDownloadXLSX`。下载请求使用 `GET https://scraperapi.dataify.com/download`，以 `api_key`、`task_id`、`type` 查询参数认证和选择格式；不会添加 Bearer Header。任务 ID 仍由服务端校验。

## 11. Web Unlocker 示例

```go
resp, err := client.WebUnlocker.Request(ctx, dataify.RequestWebUnlockerRequest{
	URL:      "https://example.com",
	Type:     "html",
	JSRender: "true",
	Country:  "us",
})
```

如果需要同时返回 HTML 和 PNG：

```go
resp, err := client.WebUnlocker.Request(ctx, dataify.RequestWebUnlockerRequest{
	URL: "https://example.com",
	Type: "html,png",
})
```

## 12. 错误处理

缺少 API Key：

```go
if errors.Is(err, dataify.ErrMissingToken) {
	log.Fatal("missing Dataify API token")
}
```

HTTP 状态码错误：

```go
var apiErr *dataify.APIError
if errors.As(err, &apiErr) {
	log.Printf("status=%d body=%s", apiErr.StatusCode, string(apiErr.Body))
}
```

完整示例：

```go
resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{
	Q:    "Dataify",
	JSON: "1",
})
if err != nil {
	var apiErr *dataify.APIError
	switch {
	case errors.Is(err, dataify.ErrMissingToken):
		log.Fatal("missing Dataify API token")
	case errors.As(err, &apiErr):
		log.Fatalf("Dataify API error: status=%d body=%s", apiErr.StatusCode, string(apiErr.Body))
	default:
		log.Fatal(err)
	}
}

fmt.Println(resp.String())
```

## 13. 如何查找工具和参数

SDK 里的工具按平台分组，例如：

- `client.Google.Search`
- `client.Bing.Search`
- `client.Yandex.Search`
- `client.Amazon.Product`
- `client.YouTube.Product`
- `client.TikTok.Posts`
- `client.WebUnlocker.Request`

每个方法都有对应的请求结构体，例如：

- `dataify.GoogleSearchRequest`
- `dataify.AmazonProductRequest`
- `dataify.RequestWebUnlockerRequest`

完整工具和参数列表见：

```text
docs/tools.md
docs/tools.yaml
```

如果你使用 IDE，例如 GoLand 或 VS Code，可以输入 `client.` 查看所有服务分组，再输入服务名查看可用方法。

## 14. 在其他项目中使用

在任意 Go 项目中：

```cmd
go mod init example.com/my-project
go get github.com/dataify-server/dataify-sdk-go@v0.1.0
```

代码中引入：

```go
import dataify "github.com/dataify-server/dataify-sdk-go"
```

然后初始化：

```go
client := dataify.NewClient(dataify.TokenFromEnv())
```

不要使用本地 `replace`，否则就不是在测试别人从 GitHub 拉取 SDK 的真实场景。

## 15. 发布后拉取测试

新建一个空目录：

```cmd
cd /d D:\codex_workspace
mkdir dataify-sdk-go-install-test
cd dataify-sdk-go-install-test
```

初始化测试项目：

```cmd
go mod init example.com/dataify-sdk-go-install-test
go get github.com/dataify-server/dataify-sdk-go@v0.1.0
```

写入最小测试代码后执行：

```cmd
go run .
```

如果可以正常编译并调用，说明其他项目可以正常使用这个 SDK。

## 16. 常见问题

### 14.1 `missing API token`

说明没有传 API Token。优先检查标准环境变量；旧变量只用于兼容已有环境：

```cmd
echo %DATAIFY_API_TOKEN%
```

PowerShell：

```powershell
echo $env:DATAIFY_API_TOKEN
```

兼容旧环境时，也可以检查 `DATAIFY_TOKEN` 或 `DATAIFY_API_KEY`；新项目应迁移到 `DATAIFY_API_TOKEN`。

### 14.2 `module declares its path as ...`

说明 GitHub 仓库路径和 `go.mod` 里的 `module` 不一致。当前 SDK 应该统一使用：

```text
github.com/dataify-server/dataify-sdk-go
```

### 14.3 `unknown revision v0.1.0`

说明远程 tag 没有推送成功，或 Go 代理还没有刷新。可以先直接指定最新提交：

```cmd
go get github.com/dataify-server/dataify-sdk-go@main
```

正式使用建议打版本 tag，例如：

```cmd
git tag v0.1.1
git push origin v0.1.1
```

然后：

```cmd
go get github.com/dataify-server/dataify-sdk-go@v0.1.1
```

### 14.4 HTTP 400

通常是业务参数不正确或缺少必填参数。检查对应请求结构体和 [tools.md](tools.md) 中的参数说明。

### 14.5 HTTP 401 或 403

通常是 API Key 不正确、额度不足、权限不足或请求头鉴权失败。确认：

- 环境变量是否设置正确。
- API Key 是否有效。
- 当前账户是否有对应工具额度。

## 17. 安全建议

- 不要把 API Key 写死在 `.go` 文件里。
- 不要提交 `.env` 文件。
- 不要在日志里打印完整 API Key。
- 不要把真实 API Key 写进 README、issue、commit message。
- CI/CD 中使用平台的 Secret 管理功能注入环境变量。
