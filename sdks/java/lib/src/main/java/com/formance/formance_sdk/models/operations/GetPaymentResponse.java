/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class GetPaymentResponse {
    public String contentType;
    public GetPaymentResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * OK
     */
    public com.formance.formance_sdk.models.shared.PaymentResponse paymentResponse;
    public GetPaymentResponse withPaymentResponse(com.formance.formance_sdk.models.shared.PaymentResponse paymentResponse) {
        this.paymentResponse = paymentResponse;
        return this;
    }
    
    public Integer statusCode;
    public GetPaymentResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public GetPaymentResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}