/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export enum ErrorErrorCodeEnum {
  Validation = "VALIDATION",
}

/**
 * General error
 */
export class ErrorT extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "errorCode" })
  errorCode: ErrorErrorCodeEnum;

  @SpeakeasyMetadata()
  @Expose({ name: "errorMessage" })
  errorMessage: string;
}