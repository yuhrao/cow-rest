# Cow Rest
![go-version](https://img.shields.io/github/go-mod/go-version/YuhriBernardes/cow-rest)

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

You can find the binaries for download on [releases page](https://github.com/YuhriBernardes/cow-rest/releases)

Enter the binary location and run.

**linux**: `./cowrest`

**MacOS**: `./darwin_cowrest`

**windows**: `./cowrest.exe`


## API documentation

You can get the API Docs on [this link](https://app.swaggerhub.com/apis/Yuhri-Bernardes/cowrest/1.0.0)

The OpenAPI specs can be found in `doc/` folder. You can use one of the tools below for visualization:

- [Postman](https://www.postman.com)
- [Insomnia Designer](https://insomnia.rest/download/)



## Next steps

- [ ] Get cow files
- [ ] Create new cows
- [ ] Generate swagger specs automatically