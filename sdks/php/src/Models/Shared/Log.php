<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


class Log
{
    /**
     * $data
     * 
     * @var array<string, mixed> $data
     */
	#[\JMS\Serializer\Annotation\SerializedName('data')]
    #[\JMS\Serializer\Annotation\Type('array<string, mixed>')]
    public array $data;
    
	#[\JMS\Serializer\Annotation\SerializedName('date')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    public \DateTime $date;
    
	#[\JMS\Serializer\Annotation\SerializedName('hash')]
    #[\JMS\Serializer\Annotation\Type('string')]
    public string $hash;
    
	#[\JMS\Serializer\Annotation\SerializedName('id')]
    #[\JMS\Serializer\Annotation\Type('int')]
    public int $id;
    
	#[\JMS\Serializer\Annotation\SerializedName('type')]
    #[\JMS\Serializer\Annotation\Type('enum<formance\stack\Models\Shared\LogTypeEnum>')]
    public LogTypeEnum $type;
    
	public function __construct()
	{
		$this->data = [];
		$this->date = new \DateTime();
		$this->hash = "";
		$this->id = 0;
		$this->type = \formance\stack\Models\Shared\LogTypeEnum::NEW_TRANSACTION;
	}
}