# NodeJS HTTP server

This package contains a basic HTTP server written in NodeJS that can be used to consume webhooks from Plain.

The server binds to `localhost` and listens on port `3060`.

## Usage

First install the dependencies. We are using [pnpm](https://pnpm.io/) to manage the dependencies, but you can also use
npm.

```bash
pnpm install
```

Then run the server:

```text
pnpm run start
```

The server requires you to use basic HTTP authentication. You can use the default credentials `username:password` or
pick your own by setting the environment variables `USERNAME` and `PASSWORD`:

```bash
USERNAME=pete PASSWORD=s3Cr3t pnpm run start
```

Remember that to use these credentials you need to include them in the HTTPS URL you get from the reverse proxy. For
instance, if you get the URL `https://example.com`, the URL to use (with the default credentials)
is `https://username:password@example.com`



To stop the server, press `Ctrl+C`.
