# CHTTP

An opinionated, chainable, elegant HTTP client. Based on net/http.

⚠**️ This project is still working in progress. Don't use it in production environment.**

## Installation

```bash
go get github.com/h4rvey99/chttp
```

## Usage

Send a GET request.

```go
res := chttp.Get("http://mockbin.org/request").ToString()
fmt.Println(res.Status())  // 200
fmt.Println(res.String())  // HTTP response body string
```

With URL parameters.

```go
res := chttp.Get("http://mockbin.org/request").
	Param("id", 1).
	Param("category", 2).
	ToString()
```

With HTTP headers.

```go
res := chttp.Get("http://mockbin.org/request").
	Meta("Authorization", "Bearer <TOKEN>").
	ToString()
```

Unmarshal JSON response body.

```go
var book struct {
	ID int        `json:"id"`
	Title string  `json:"title"`
	Author string `json:"author"`
}
res := chttp.Get("http://localhost:8080/book").ToStruct(&book)
if res.Error() != nil {
	log.Fatal("can't get book", err)
}
fmt.Printf("%+v", book)
```

