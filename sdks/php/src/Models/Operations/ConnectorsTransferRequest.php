<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;

use \formance\stack\Utils\SpeakeasyMetadata;
class ConnectorsTransferRequest
{
	#[SpeakeasyMetadata('request:mediaType=application/json')]
    public \formance\stack\Models\Shared\TransferRequest $transferRequest;
    
    /**
     * The name of the connector.
     * 
     * @var \formance\stack\Models\Shared\ConnectorEnum $connector
     */
	#[SpeakeasyMetadata('pathParam:style=simple,explode=false,name=connector')]
    public \formance\stack\Models\Shared\ConnectorEnum $connector;
    
	public function __construct()
	{
		$this->transferRequest = new \formance\stack\Models\Shared\TransferRequest();
		$this->connector = \formance\stack\Models\Shared\ConnectorEnum::STRIPE;
	}
}
