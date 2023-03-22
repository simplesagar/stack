/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { ErrorsEnumEnum } from "./errorsenumenum";
import { Transaction } from "./transaction";
import { Expose, Type } from "class-transformer";

/**
 * On success, it will return a 200 status code, and the resulting transaction under the `transaction` field.
 *
 * @remarks
 *
 * On failure, it will also return a 200 status code, and the following fields:
 *   - `details`: contains a URL. When there is an error parsing Numscript, the result can be difficult to read—the provided URL will render the error in an easy-to-read format.
 *   - `errorCode` and `error_code` (deprecated): contains the string code of the error
 *   - `errorMessage` and `error_message` (deprecated): contains a human-readable indication of what went wrong, for example that an account had insufficient funds, or that there was an error in the provided Numscript.
 *
 */
export class ScriptResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "details" })
  details?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "errorCode" })
  errorCode?: ErrorsEnumEnum;

  @SpeakeasyMetadata()
  @Expose({ name: "errorMessage" })
  errorMessage?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "transaction" })
  @Type(() => Transaction)
  transaction?: Transaction;
}
