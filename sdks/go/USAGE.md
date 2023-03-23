<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/formancehq/formance-sdk-go"
    "github.com/formancehq/formance-sdk-go/pkg/models/shared"
    "github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    req := operations.AddScopeToClientRequest{
        ClientID: "unde",
        ScopeID: "deserunt",
    }

    ctx := context.Background()
    res, err := s.Auth.AddScopeToClient(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->