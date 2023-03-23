<!-- Start SDK Example Usage -->
```java
package hello.world;

import com.formance.formance_sdk.SDK;
import com.formance.formance_sdk.models.shared.Security;
import com.formance.formance_sdk.models.operations.AddScopeToClientRequest;
import com.formance.formance_sdk.models.operations.AddScopeToClientResponse;

public class Application {
    public static void main(String[] args) {
        try {
            SDK sdk = SDK.builder()
                .setSecurity(new Security() {{
                    authorization = "Bearer YOUR_ACCESS_TOKEN_HERE";
                }})
                .build();

            AddScopeToClientRequest req = new AddScopeToClientRequest() {{
                clientId = "unde";
                scopeId = "deserunt";
            }}            

            AddScopeToClientResponse res = sdk.auth.addScopeToClient(req);

            if (res.statusCode == 200) {
                // handle response
            }
        } catch (Exception e) {
            // handle exception
        }
```
<!-- End SDK Example Usage -->