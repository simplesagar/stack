/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";
import { Type } from "class-transformer";

export class GetHoldRequest extends SpeakeasyBase {
  /**
   * The hold ID
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=holdID",
  })
  holdID: string;
}

export class GetHoldResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  /**
   * Holds
   */
  @SpeakeasyMetadata()
  getHoldResponse?: shared.GetHoldResponse;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;

  /**
   * Error
   */
  @SpeakeasyMetadata()
  walletsErrorResponse?: shared.WalletsErrorResponse;
}
