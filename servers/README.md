# Servers

This directory contains a collection of sample HTTP servers you can use to consume webhooks from Plain.

Plain only supports HTTPS webhook endpoint URLs. The servers included here are bound to `localhost`, so you will need a tool like `localtunnel` or `ngrok` to expose your server to the internet.

Pick any of the following tools and follow the instructions below to get started.

## localtunnel

First, install `localtunnel` globally:

```bash
npm install -g localtunnel
```

Then, start it as a reverse proxy for the HTTP server:

```bash
lt --port 3060
```

This will give you an HTTPS URL that you can use to configure your webhook target in Plain's Support App.

## ngrok

First, install `ngrok` following the instructions [here](https://ngrok.com/download). You might need to sign up for an account.

Once ngrok is installed, run it as a reverse proxy for the HTTP server:

```bash
ngrok http localhost:3060
```
