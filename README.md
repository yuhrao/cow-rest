# Cow Rest

A REST API interface to use [ `cowsay` ](https://github.com/schacon/cowsay)
- [Cow Rest](#cow-rest)
  - [Start application](#start-application)
    - [With Go](#with-go)
    - [Using the binary](#using-the-binary)
  - [API documentation](#api-documentation)
  - [Next steps](#next-steps)

## Start application

### With Go

> Remember to configure your GOPATH correctly

```shell
go run ./main.go
```

To build the app, run `go build -o cowrest ./main.go`

### Using the binary

Enter the binary location and run

**linux**: `./cowrest`

**MacOS**: `./darwin_cowrest`

**windows**: `./cowrest.exe`


## API documentation

You can find an OpenAPI spec in `doc/` folder. You can use one of the tools below:

- [Postman](https://www.postman.com)
- [Insomnia Designer](https://insomnia.rest/download/)
- [SwaggerUI](https://app.swaggerhub.com/apis/Yuhri-Bernardes/cowrest/1.0.0) (only need your browser)


## Next steps

- [ ] Get cow files
- [ ] Create new cows
- [ ] Generate swagger specs automatically