### Go client for Posthook API

Implements the API documented at [https://docs.posthook.io/](https://docs.posthook.io/)

Install with `go get -u github.com/pims/posthook`

example use:

```go
// create posthook client
ph := posthook.New(os.Getenv("POSTHOOK_API_KEY"))

// verify the endpoint being used
fmt.Println(ph.Endpoint())

// data to be sent back to us when posthook executes the hook
data := map[string]interface{}{
    "foobar": "baz",
}

// Step 1: schedule a hook
hook, err := ph.Schedule("ping", time.Now().Add(1*time.Minute).UTC(), data)
handleErr(err)
```
