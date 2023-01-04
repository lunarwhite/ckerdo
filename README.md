# load-balancer

![GitHub top language](https://img.shields.io/github/languages/top/lunarwhite/load-balancer) [![Go Report Card](https://goreportcard.com/badge/github.com/lunarwhite/load-balancer)](https://goreportcard.com/report/github.com/lunarwhite/load-balancer) [![codecov](https://codecov.io/gh/lunarwhite/load-balancer/branch/master/graph/badge.svg)](https://codecov.io/gh/lunarwhite/load-balancer) [![GitHub Action](https://github.com/lunarwhite/load-balancer/actions/workflows/go.yml/badge.svg)](https://github.com/lunarwhite/load-balancer/actions/workflows/go.yml) [![Releases](https://img.shields.io/github/release/lunarwhite/load-balancer/all.svg)](https://github.com/lunarwhite/load-balancer/releases)

## Introduction

This project is a simple layer 7 load balancer written in Go, supports both http and https scheme, and implements several common load balancing algorithms. Note that not stable for production use, just for personal practice.

## Highlights

It currently supports `round-robin`, `random`, `consistent-hash`, `ip-hash` and `least-load` algorithms

## Getting Started

## Contributing

Contributions are welcome. If you are open source newcomer, you can follow [this guide](https://opensource.guide/how-to-contribute/) by GitHub.

## Report Vulnerability

If you come across a security related issue, please open an issue.

## License

This project is licensed under the term of the [MIT License](https://github.com/lunarwhite/load-balancer/blob/main/LICENSE).