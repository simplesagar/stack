/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * PaymentResponse - OK
 */
public class PaymentResponse {
    @JsonProperty("data")public Payment data;
    public PaymentResponse withData(Payment data) {
        this.data = data;
        return this;
    }
    
}