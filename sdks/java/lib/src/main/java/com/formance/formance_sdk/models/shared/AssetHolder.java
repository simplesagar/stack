/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class AssetHolder {
    @JsonProperty("assets")public java.util.Map<String, Double> assets;
    public AssetHolder withAssets(java.util.Map<String, Double> assets) {
        this.assets = assets;
        return this;
    }
    
}