<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class WorkflowInstanceHistory
{
	#[\JMS\Serializer\Annotation\SerializedName('error')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $error = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('input')]
    #[\JMS\Serializer\Annotation\Type('mixed')]
    public mixed $input;
    
	#[\JMS\Serializer\Annotation\SerializedName('name')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $name;
    
	#[\JMS\Serializer\Annotation\SerializedName('startedAt')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    public \DateTime $startedAt;
    
	#[\JMS\Serializer\Annotation\SerializedName('terminated')]
    #[\JMS\Serializer\Annotation\Type('bool')]
    public bool $terminated;
    
	#[\JMS\Serializer\Annotation\SerializedName('terminatedAt')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?\DateTime $terminatedAt = null;
    
	public function __construct()
	{
		$this->error = null;
		$this->input = null;
		$this->name = "";
		$this->startedAt = new \DateTime();
		$this->terminated = false;
		$this->terminatedAt = null;
	}
}
