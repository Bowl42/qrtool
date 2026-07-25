# qrtool

Small HTTP service that generates QR Code images from GET requests.

## Run

```sh
go run ./cmd/qrtool
```

The service listens on `:8080` by default. Set `PORT` to override it.

## Docker

```sh
docker build -t qrtool .
docker run --rm -p 8080:8080 qrtool
```

## API

```text
GET /?text=hello
GET /anything.png?text=hello
GET /anything.svg?text=hello
GET /healthz
```

The path is optional. When present, the file name is used for browser save-as
behavior. Responses are served with `Content-Disposition: inline`, so opening
the URL directly shows the image in the browser.

Query parameters:

| Name | Default | Description |
| --- | --- | --- |
| `text` | required | Text encoded into the QR Code |
| `type` | `qr` | Code type. Only `qr` is supported in this version |
| `size` | `256` | PNG size in pixels, from `64` to `2048` |
| `level` | `m` | Error correction level: `l`, `m`, `q`, or `h` |
| `margin` | `4` | Quiet-zone modules, from `0` to `20` |

Supported output formats:

- `png`
- `svg`

Examples:

```sh
curl 'http://localhost:8080/?text=hello' --output qrcode.png
curl 'http://localhost:8080/invite.svg?text=https://example.com' --output invite.svg
```

## License

MIT
