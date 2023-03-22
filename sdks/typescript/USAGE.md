<!-- Start SDK Example Usage -->
```typescript
import {
  GetServerInfoResponse
} from "@formance/formance-sdk/dist/sdk/models/operations";

import { AxiosError } from "axios";
import { SDK } from "@formance/formance-sdk";
const sdk = new SDK({
  security: {
    authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
  },
});

sdk.getServerInfo().then((res: GetServerInfoResponse | AxiosError) => {
   // handle response
});
```
<!-- End SDK Example Usage -->