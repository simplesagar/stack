/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class ChangeConfigSecretResponse {
    /**
     * Secret successfully changed.
     */
    public com.formance.formance_sdk.models.shared.ConfigResponse configResponse;
    public ChangeConfigSecretResponse withConfigResponse(com.formance.formance_sdk.models.shared.ConfigResponse configResponse) {
        this.configResponse = configResponse;
        return this;
    }
    
    public String contentType;
    public ChangeConfigSecretResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    public Integer statusCode;
    public ChangeConfigSecretResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public ChangeConfigSecretResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}
