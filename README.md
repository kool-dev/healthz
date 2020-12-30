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

```bash
healthz -i '[{"name": "check 1", "type": "tcp", "value": "localhost:80"}, {"name": "check 2", "type": "http", "value": "http://localhost"}, {"name": "check 3", "type": "exec", "value": "ls -lah /"}]'
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
