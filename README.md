# ArtisanDB (In Development)

ArtisanDB is an in-memory key-value database built completely with Go.

## Setup

ArtisanDB requires pretty much no setup. The only way to get ArtisanDB right now is to build it from this repository.

If you wish to use ArtisanDB, clone this repository and run `go build main.go` then `./main`.

The database server will be started and will be available at `localhost:3333`

## Usage

ArtisanDB currently supports 5 major operations. `GET`, `SET`, `COUNT`, `ALL`

ArtisanDB exposes its API through a simple TCP server which accepts simple messages.

An ArtisanDB message conforms to the following format:

```
artisan:{{id}}
{{username}}:{{password}}
{{method}} {{key}}
{{value}}
```


An ArtisanDB message for a SET operation with the key "city" and value "London" sent by a user with username "admin" and password "adminpassword" will be formatted as follows:
```
artisan:765467897654
admin:adminpassword
SET city
London
```

Artisan messages are terminated with a Carriage Return, Line Feed, Carriage Return, Line Feed `\r\n\r\n` sequence.

## Coming soon

ArtisanDB has pending features on the roadmap and in active development. In order of priority, these are the features in active development:

- TLS support
- Persistent storage even after restarts
- TTL support for stored data
- A CLI tool for managing configurations
