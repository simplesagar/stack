/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * StatsResponse - OK
 */
public class StatsResponse {
    @JsonProperty("data")public Stats data;
    public StatsResponse withData(Stats data) {
        this.data = data;
        return this;
    }
    
}
