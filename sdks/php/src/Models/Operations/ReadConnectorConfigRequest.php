<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;

use \formance\stack\Utils\SpeakeasyMetadata;
class ReadConnectorConfigRequest
{
    /**
     * The name of the connector.
     * 
     * @var \formance\stack\Models\Shared\ConnectorEnum $connector
     */
	#[SpeakeasyMetadata('pathParam:style=simple,explode=false,name=connector')]
    public \formance\stack\Models\Shared\ConnectorEnum $connector;
    
	public function __construct()
	{
		$this->connector = \formance\stack\Models\Shared\ConnectorEnum::STRIPE;
	}
}
