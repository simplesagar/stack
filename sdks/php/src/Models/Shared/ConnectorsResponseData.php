<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class ConnectorsResponseData
{
	#[\JMS\Serializer\Annotation\SerializedName('enabled')]
    #[\JMS\Serializer\Annotation\Type('bool')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?bool $enabled = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('provider')]
    #[\JMS\Serializer\Annotation\Type('enum<formance\stack\Models\Shared\ConnectorEnum>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?ConnectorEnum $provider = null;
    
	public function __construct()
	{
		$this->enabled = null;
		$this->provider = null;
	}
}
