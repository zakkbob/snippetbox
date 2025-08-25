# TLS

The server is setup to serve over TLS, so a certificate must be present in this
directory for the server to run.

For local development and testing, one can be generated using `generate_cert.go`.

For example (location may vary):
`go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`
