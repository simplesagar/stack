/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class UninstallConnectorRequest {
    /**
     * The name of the connector.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=connector")public com.formance.formance_sdk.models.shared.ConnectorEnum connector;
    public UninstallConnectorRequest withConnector(com.formance.formance_sdk.models.shared.ConnectorEnum connector) {
        this.connector = connector;
        return this;
    }
    
}
