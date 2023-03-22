/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * TransfersResponse - OK
 */
public class TransfersResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")public TransfersResponseData[] data;
    public TransfersResponse withData(TransfersResponseData[] data) {
        this.data = data;
        return this;
    }
    
}