# Girder
**A Go web API toolkit designed to reduce common boilerplate**

**⚠ WARNING** Opinionated code ahead

 Girder is a relatively opinionated toolkit, it assumes you will be building
 an API which consumes structured, serializable objects (JSON for example) and
 then produces responses in a similar format.

 It has been inspired from a wide number of different projects I've developed over
 the years on everything from C# to Node.js and finally Go, so it'll take inspiration
 from some of those along the way. That being said, Girder doesn't aim to be a
 framework upon which you'll build a full application - instead it is meant to provide
 just the glue between your handlers and the `http` layer.

## Example

 ```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/SierraSoftworks/girder"
    "github.com/gorilla/mux"
)

func hello(c *girder.Context) (interface{}, error) {
    return fmt.Sprintf("Hello %s", c.Vars["name"]), nil
}

func main() {
    h := girder.NewHandler(hello)

    r := mux.NewRouter()
    r.Path("/api/v1/hello/{name}").Method("GET").Handler(h)

    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
```

## Design
Girder is designed such that all requests are routed through a `girder.Handler`. This handler
is responsible for all request pre-processing, dispatching the request and its context to your
handler function, and finally serializing the response from your handler function into the response.

In addition to this, it is responsible for converting any errors you return into a well formatted
error object.

A `girder.HandlerFunc` looks like this, keeping things nice and simple to write and with very little
overhead.

```go
func handlerFunc(c *girder.Context) (interface{}, error) {
    return MyData{"x"}, nil
}
```

Girder's handlers have a `Preprocessors` slice which contains functions which will be executed
before your handler. These functions may optionally return an error to bypass further execution
and are an excellent place to handle things like authentication and authorization or logging.

## Usage

### Deserializing a Request
You'll commonly implement systems in which you accept data as part of a `POST` request's body.
In these cases, you'll want an easy way to deserialize the request into some object of your
choosing. Girder makes this spectacularly simple, here's a quick echo function.

```go
func handlerFunc(c *girder.Context) (interface{}, error) {
    var req RequestData
    if err := c.ReadBody(&req); err != nil {
        return nil, err
    }

    return req, nil
}
```

### Accessing Gorilla Mux Route Parameters
Girder expects that you, like every other Go developer, will be using [Gorilla Mux](https://github.com/gorilla/mux)
as your router of choice. One of the great features it provides out of the box is support for route
parameters. You can access your route parameters directly from the context using the `Vars` property.

```go
// GET /api/v1/hello/{name}
func handlerFunc(c *girder.Context) (interface{}, error) {
    name := c.Vars["name"]
    return fmt.Sprintf("Hello %s", name), nil
}
```

### User Authentication
User authentication is a common enough use case that we've decided to build it into Girder
out of the box. Users are provided by a callback when you register your authorization preprocessor
and are expected to provide both a `GetID()` method (for use within your application) and a
`HasPermission()` method which allows Girder to determine whether the user has permission to
access a route or not.

```go
package main

import (
    "net/http"

    "github.com/SierraSoftworks/girder"
    "github.com/SierraSoftworks/girder/errors"
    "github.com/gorilla/mux"
)

type UserStore struct {
    users []User
}

func (s *UserStore) GetUser(token *girder.AuthorizationToken) (girder.User, error) {
    if token.Type != "Token" {
        return nil, errors.NewError(401, "Unauthorized", "You failed to provide a valid authentication token type with your request.")
    }

    for _, user := range s.users {
        for _, t := range user.tokens {
            if t == token.Value {
                return user, nil
            }
        }
    }

    return nil, nil
}

type User struct {
    id         string
    tokens      []string
}

// Extend your user type with the GetID() and HasPermission() functions
func (u *User) GetID() string {
    return u.id
}

func (u *User) HasPermission(permission string) bool {
    return true
}

// GET /api/v1/hello
func hello(c *girder.Context) (interface{}, error) {
    user, err := users.GetByID(c.User.ID())
    if err != nil {
        return nil, errors.From(err)
    }

    return fmt.Sprintf("Hello %s", user.Name), nil
}

func main() {
    // Point girder at your user store
    store := &UserStore{
        users: []User{
            User{
                id: "bob",
                tokens: []string{"0123456789abcdef"},
            },
        },
    }

    r := mux.NewRouter()
    r.Path("/api/v1/hello").Handler(girder.NewHandler(hello).RequireAuthentication(store.GetUser).RequirePermission("hello"))

    http.ListenAndServe(":8080", r)
}
```