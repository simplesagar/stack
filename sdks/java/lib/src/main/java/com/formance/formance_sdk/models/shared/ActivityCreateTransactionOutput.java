/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityCreateTransactionOutput {
    @JsonProperty("data")public Transaction[] data;
    public ActivityCreateTransactionOutput withData(Transaction[] data) {
        this.data = data;
        return this;
    }
    
}
