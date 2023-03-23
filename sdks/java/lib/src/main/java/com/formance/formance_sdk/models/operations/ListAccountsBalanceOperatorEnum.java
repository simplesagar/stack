/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonValue;

/**
 * ListAccountsBalanceOperatorEnum - Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
 * 
 */
public enum ListAccountsBalanceOperatorEnum {
    GTE("gte"),
    LTE("lte"),
    GT("gt"),
    LT("lt"),
    E("e"),
    NE("ne");

    @JsonValue
    public final String value;

    private ListAccountsBalanceOperatorEnum(String value) {
        this.value = value;
    }
}
