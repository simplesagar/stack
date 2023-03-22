<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * ScriptResponse - On success, it will return a 200 status code, and the resulting transaction under the `transaction` field.
 * 
 * 
 * On failure, it will also return a 200 status code, and the following fields:
 *   - `details`: contains a URL. When there is an error parsing Numscript, the result can be difficult to read—the provided URL will render the error in an easy-to-read format.
 *   - `errorCode` and `error_code` (deprecated): contains the string code of the error
 *   - `errorMessage` and `error_message` (deprecated): contains a human-readable indication of what went wrong, for example that an account had insufficient funds, or that there was an error in the provided Numscript.
 * 
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class ScriptResponse
{
	#[\JMS\Serializer\Annotation\SerializedName('details')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $details = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('errorCode')]
    #[\JMS\Serializer\Annotation\Type('enum<formance\stack\Models\Shared\ErrorsEnumEnum>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?ErrorsEnumEnum $errorCode = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('errorMessage')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $errorMessage = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('transaction')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\Transaction')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?Transaction $transaction = null;
    
	public function __construct()
	{
		$this->details = null;
		$this->errorCode = null;
		$this->errorMessage = null;
		$this->transaction = null;
	}
}