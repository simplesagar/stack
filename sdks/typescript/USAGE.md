<!-- Start SDK Example Usage -->
```typescript
import {
  AddScopeToClientRequest,
  AddScopeToClientResponse
} from "@formance/formance-sdk/dist/sdk/models/operations";

import { AxiosError } from "axios";
import { SDK } from "@formance/formance-sdk";
const sdk = new SDK({
  security: {
    authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
  },
});

const req: AddScopeToClientRequest = {
  clientId: "unde",
  scopeId: "deserunt",
};

sdk.auth.addScopeToClient(req).then((res: AddScopeToClientResponse | AxiosError) => {
   // handle response
});
```
<!-- End SDK Example Usage -->