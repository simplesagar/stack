/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class ListUsersResponse {
    public String contentType;
    public ListUsersResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * List of users
     */
    public com.formance.formance_sdk.models.shared.ListUsersResponse listUsersResponse;
    public ListUsersResponse withListUsersResponse(com.formance.formance_sdk.models.shared.ListUsersResponse listUsersResponse) {
        this.listUsersResponse = listUsersResponse;
        return this;
    }
    
    public Integer statusCode;
    public ListUsersResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public ListUsersResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}