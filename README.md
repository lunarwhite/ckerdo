# load-balancer

![GitHub top language](https://img.shields.io/github/languages/top/lunarwhite/load-balancer) [![Go Report Card](https://goreportcard.com/badge/github.com/lunarwhite/load-balancer)](https://goreportcard.com/report/github.com/lunarwhite/load-balancer) [![codecov](https://codecov.io/gh/lunarwhite/load-balancer/branch/master/graph/badge.svg)](https://codecov.io/gh/lunarwhite/load-balancer) [![GitHub Action](https://github.com/lunarwhite/load-balancer/actions/workflows/go.yml/badge.svg)](https://github.com/lunarwhite/load-balancer/actions/workflows/go.yml) [![Releases](https://img.shields.io/github/release/lunarwhite/load-balancer/all.svg)](https://github.com/lunarwhite/load-balancer/releases)

## Introduction

This project is a simple layer 7 load balancer written in Go, supports both http and https scheme, and implements several common load balancing algorithms. Note that not stable for production use, just for personal practice.

## Highlights

It currently supports `round-robin`, `random`, `consistent-hash`, `ip-hash` and `least-load` algorithms.

It will perform `health check` on all proxy sites periodically. When the site is unreachable, it will be removed from the balancer automatically . However, `load-balancer` will still perform `health check` on unreachable sites. When the site is reachable, it will add it to the balancer automatically.

## Getting Started

### Build and run

Clone the source code:

```shell
git clone https://github.com/lunarwhite/load-balancer.git
```

You should specific the `config.yaml` file first, example see [config.yaml](./examples/config.yaml):

```yaml
schema: http # support `http` and `https`
port: 8089 # port for balancer
location: # route matching for reverse proxy
  - pattern: /
    proxy_pass: # URLs of the reverse proxy
      - "https://192.168.1.1"
      - "https://192.168.1.2"
      - "https://192.168.1.3"
      - "https://192.168.1.4"
    balance_algo: round-robin # supprt `round-robin`,`random`, `ip-hash, `consistent-hash`, `least-load`
ssl_certificate: # your ssl certificate
ssl_certificate_key: # your ssl certificate key
max_req_allowed: 100 # The max number of requests that the balancer can handle at the same time, 0 refers to no limit
enable_health_check: true
health_check_interval: 3 # second
```

Then, run commands to build:

```shell
cd ./load-balancer
go build
```

Execute `load-balancer`, it will print the configuration details:

```shell
$ ./load-balancer

Schema: http
Port: 8089
Location:
        Route: /
        Proxy Pass: [https://192.168.1.1 https://192.168.1.2 https://192.168.1.3 https://192.168.1.4]
        Algo: round-robin
Enable Health Check: true
```

### API Usage

You can use it as a third-party Go lib in your project.

```shell
go get github.com/lunarwhite/load-balancer/balancer
```

Build the load balancer with `balancer.Build`:

```go
hosts := []string{
	"http://192.168.10.102",
	"http://192.168.10.103",
	"http://192.168.10.104",
	"http://192.168.10.105",
}

lb, err := balancer.Build(balancer.LeastLoadBalancer, hosts)
if err != nil {
	return err
}
```

Then, you can use `load-balancer` like this:

```go
clientAddr := "172.160.10.2"  // request IP

targetHost, err := lb.Balance(clientAddr)
if err != nil {
	log.Fatal(err)
}

lb.Inc(targetHost)
defer lb.Done(targetHost)

// route to target host
```

Each load balancer implements the `balancer.Balancer` interface:

```go
type Balancer interface {
	Add(string)
	Remove(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}
```

## Contributing

Contributions are welcome. If you are open source newcomer, you can follow [this guide](https://opensource.guide/how-to-contribute/) by GitHub.

## Report Vulnerability

If you come across a security related issue, please open an issue.

## License

This project is licensed under the term of the [MIT License](https://github.com/lunarwhite/load-balancer/blob/main/LICENSE).
