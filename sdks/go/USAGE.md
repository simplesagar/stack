<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/formancehq/formance/sdks/go"
    "github.com/formancehq/formance/sdks/go/pkg/models/shared"
    "github.com/formancehq/formance/sdks/go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.GetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->