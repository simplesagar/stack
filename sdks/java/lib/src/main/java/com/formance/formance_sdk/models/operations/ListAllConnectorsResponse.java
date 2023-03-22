/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class ListAllConnectorsResponse {
    /**
     * OK
     */
    public com.formance.formance_sdk.models.shared.ConnectorsResponse connectorsResponse;
    public ListAllConnectorsResponse withConnectorsResponse(com.formance.formance_sdk.models.shared.ConnectorsResponse connectorsResponse) {
        this.connectorsResponse = connectorsResponse;
        return this;
    }
    
    public String contentType;
    public ListAllConnectorsResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    public Integer statusCode;
    public ListAllConnectorsResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public ListAllConnectorsResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}