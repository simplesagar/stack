<!-- Start SDK Example Usage -->
```python
import sdk
from sdk.models import operations, shared

s = sdk.SDK(
    security=shared.Security(
        authorization="Bearer YOUR_ACCESS_TOKEN_HERE",
    ),
)

    
res = s.get_server_info()

if res.server_info is not None:
    # handle response
```
<!-- End SDK Example Usage -->