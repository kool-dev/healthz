# healthz

[![Kool.dev](https://kool.dev/img/logo.png)](https://kool.dev)

Utility to simplify health check on container applications.

---

[![Go Report Card](https://goreportcard.com/badge/github.com/kool-dev/healthz)](https://goreportcard.com/report/github.com/kool-dev/healthz)
[![codecov](https://codecov.io/gh/kool-dev/healthz/branch/main/graph/badge.svg?token=6JP8JRVF5U)](https://codecov.io/gh/kool-dev/healthz)
[![Maintainability](https://api.codeclimate.com/v1/badges/8cbfe7f40b29386e532d/maintainability)](https://codeclimate.com/github/kool-dev/healthz/maintainability)
[<img src="https://img.shields.io/badge/Join%20Slack-kool--dev-orange?logo=slack">](https://join.slack.com/t/kool-dev/shared_invite/zt-jeh36s5g-kVFHUsyLjFENUUH4ucrxPg)

### Installation

`healthz` is available with a one-line install script for Linux:

```bash
curl -fsSL https://raw.githubusercontent.com/kool-dev/healthz/main/install.sh | bash
```

### Usage

`healthz` receives a JSON array as parameter, which contains all tests it needs to perform.

After executing all the checks, it will exit with with `0` exit code for success, and non-zero otherwise.

```bash
# check for a TCP listening port to accept connections
healthz -i '[{"type": "tcp", "value": "localhost:80"}]'

# check for an HTTP server to respond with a 200 status code
healthz -i '[{"type": "http", "value": "http://localhost"}]'

# execute a command and check for exit code to be zero
healthz -i '[{"type": "exec", "value": "ls -lah /"}]'
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
