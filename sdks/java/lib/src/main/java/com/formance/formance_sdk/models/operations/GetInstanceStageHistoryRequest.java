/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class GetInstanceStageHistoryRequest {
    /**
     * The instance id
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=instanceID")public String instanceID;
    public GetInstanceStageHistoryRequest withInstanceID(String instanceID) {
        this.instanceID = instanceID;
        return this;
    }
    
    /**
     * The stage number
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=number")public Long number;
    public GetInstanceStageHistoryRequest withNumber(Long number) {
        this.number = number;
        return this;
    }
    
}
