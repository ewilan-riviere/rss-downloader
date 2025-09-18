# slugifier

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

A small tool, written in Go, to easily download Podcast episodes from RSS feeds.

## Install

```bash
go install github.com/ewilan-riviere/rss-downloader@latest
```

## Usage

To only print episode count

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml
```

To print episodes list

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -p
```

To download episodes

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -d
```

To reverse the order of episodes

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -r
```

To limit the number of episodes processed

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -l 5
```

To download episodes to a specific directory

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -d -o path/to/dir
```

To print list as JSON file

```bash
rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml -j path/to/file.json
```

## Build

Build the script.

```bash
go build -o rss-downloader
```

You can use `./slugifier` to run the script.

```bash
./rss-downloader https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml
```

Or you can install it globally.

```bash
go install
```

## Test

Check with `curl` if the webhook is working.

```bash
go test
```

```bash
go test ./pkg/... -coverprofile=coverage.out
go test -v ./...
go test -v ./pkg/file
```

Direct usage

```bash
go run main.go https://feeds.audiomeans.fr/feed/c15bf308-42a5-4cdc-9e89-38dce9113c6b.xml
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/slugifier/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://packagist.org/packages/ewilan-riviere/slugifier
[license-src]: https://img.shields.io/github/license/ewilan-riviere/slugifier.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/slugifier/blob/main/LICENSE
