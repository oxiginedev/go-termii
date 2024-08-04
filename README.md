# go-termii

go-termii is a Go client library for accessing [Termii API](https://termii.com.ng)

## Installation

This package can be installed using:

```bash
go get github.com/oxiginedev/go-termii
```

## Usage

Import the package using:

```go
import "github.com/oxiginedev/go-termii
```

Create a new Termii client:

```go
c, err := termii.New()
if err != nil {
    // handle error
}
```
