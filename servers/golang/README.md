# GO HTTP server

This module contains a basic HTTP server written in GO that can be used to consume webhooks from Plain. It will print
out all webhook requests it receives.

The server binds to `localhost` and listens on port `3060`.

## Usage

To run the server:

```bash
go run main.go
```

The server requires you to use basic HTTP authentication. You can use the default credentials `username:password` or
pick your own by setting the environment variables `USERNAME` and `PASSWORD`:

```bash
USERNAME=pete PASSWORD=s3Cr3t go run main.go
```

Remember that to use these credentials you need to include them in the HTTPS URL you get from the reverse proxy. For
instance, if you get the URL `https://example.com`, the URL to use (with the default credentials)
is `https://username:password@example.com`

To stop the server, press `Ctrl+C`.
