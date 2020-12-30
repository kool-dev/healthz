## healthz

Utility to simplify health check applications.

### Usage

```bash
healthz -i '[{"name": "check 1", "type": "tcp", "value": "localhost:80"}, {"name": "check 2", "type": "http", "value": "http://localhost"}, {"name": "check 3", "type": "exec", "value": "ls -lah /"}]'
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
