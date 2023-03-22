<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;

use \formance\stack\Utils\SpeakeasyMetadata;
class UpdateScopeRequest
{
	#[SpeakeasyMetadata('request:mediaType=application/json')]
    public ?\formance\stack\Models\Shared\UpdateScopeRequest $updateScopeRequest = null;
    
    /**
     * Scope ID
     * 
     * @var string $scopeId
     */
	#[SpeakeasyMetadata('pathParam:style=simple,explode=false,name=scopeId')]
    public string $scopeId;
    
	public function __construct()
	{
		$this->updateScopeRequest = null;
		$this->scopeId = "";
	}
}
