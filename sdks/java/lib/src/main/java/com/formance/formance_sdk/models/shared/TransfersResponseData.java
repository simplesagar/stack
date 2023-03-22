/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class TransfersResponseData {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("amount")public Long amount;
    public TransfersResponseData withAmount(Long amount) {
        this.amount = amount;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("asset")public String asset;
    public TransfersResponseData withAsset(String asset) {
        this.asset = asset;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("currency")public String currency;
    public TransfersResponseData withCurrency(String currency) {
        this.currency = currency;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("destination")public String destination;
    public TransfersResponseData withDestination(String destination) {
        this.destination = destination;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("error")public String error;
    public TransfersResponseData withError(String error) {
        this.error = error;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("id")public String id;
    public TransfersResponseData withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("source")public String source;
    public TransfersResponseData withSource(String source) {
        this.source = source;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("status")public String status;
    public TransfersResponseData withStatus(String status) {
        this.status = status;
        return this;
    }
    
}
