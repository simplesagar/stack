/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class ListInstancesRequest {
    /**
     * Filter running instances
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=running")public Boolean running;
    public ListInstancesRequest withRunning(Boolean running) {
        this.running = running;
        return this;
    }
    
    /**
     * A workflow id
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=workflowID")public String workflowID;
    public ListInstancesRequest withWorkflowID(String workflowID) {
        this.workflowID = workflowID;
        return this;
    }
    
}
