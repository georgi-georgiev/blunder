# blunder

[![GoDoc][doc-img]][doc] [![Build][ci-img]][ci] [![GoReport][report-img]][report] [![Coverage Status][cov-img]][cov]

Package `blunder` is an error handling library with readable stack traces and flexible formatting support.

`go get github.com/georgi-georgiev/blunder`

## Install

`go install github.com/georgi-georgiev/blunder/cmd/blunder@latest`

Run `blunder gen` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files `blunder.html`.

## Swagger

When you use swagger you need to add `--parseVendor` becuase searching in `vendor` folder is disabled by default e.g. `swag init --parseVendor`

## Contributing

If you'd like to contribute to `blunder`, we'd love your input! Please submit an issue first so we can discuss your proposal.

-------------------------------------------------------------------------------

Released under the [MIT License].

[MIT License]: LICENSE.txt
[doc-img]: https://pkg.go.dev/badge/github.com/georgi-georgiev/blunder
[doc]: https://pkg.go.dev/github.com/georgi-georgiev/blunder
[ci-img]: https://github.com/georgi-georgiev/blunder/workflows/build/badge.svg
[ci]: https://github.com/georgi-georgiev/blunder/actions
[report-img]: https://goreportcard.com/badge/github.com/georgi-georgiev/blunder
[report]: https://goreportcard.com/report/github.com/georgi-georgiev/blunder
[cov-img]: https://codecov.io/gh/georgi-georgiev/blunder/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/georgi-georgiev/blunder