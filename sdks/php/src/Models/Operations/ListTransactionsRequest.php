<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;

use \formance\stack\Utils\SpeakeasyMetadata;
class ListTransactionsRequest
{
    /**
     * Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
     * 
     * @var ?string $account
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=account')]
    public ?string $account = null;
    
    /**
     * Pagination cursor, will return transactions after given txid (in descending order).
     * 
     * @var ?string $after
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=after')]
    public ?string $after = null;
    
    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     * 
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     * 
     * 
     * @var ?string $cursor
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=cursor')]
    public ?string $cursor = null;
    
    /**
     * Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
     * 
     * @var ?string $destination
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=destination')]
    public ?string $destination = null;
    
    /**
     * Filter transactions that occurred before this timestamp.
     * 
     * The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
     * 
     * 
     * @var ?\DateTime $endTime
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=endTime,dateTimeFormat=Y-m-d\TH:i:s.up')]
    public ?\DateTime $endTime = null;
    
    /**
     * Filter transactions that occurred before this timestamp.
     * 
     * The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
     * Deprecated, please use `endTime` instead.
     * 
     * 
     * @var ?\DateTime $endTimeDeprecated
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=end_time,dateTimeFormat=Y-m-d\TH:i:s.up')]
    public ?\DateTime $endTimeDeprecated = null;
    
    /**
     * Name of the ledger.
     * 
     * @var string $ledger
     */
	#[SpeakeasyMetadata('pathParam:style=simple,explode=false,name=ledger')]
    public string $ledger;
    
    /**
     * Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
     * 
     * @var ?array<string, mixed> $metadata
     */
	#[SpeakeasyMetadata('queryParam:style=deepObject,explode=true,name=metadata')]
    public ?array $metadata = null;
    
    /**
     * The maximum number of results to return per page.
     * 
     * 
     * 
     * @var ?int $pageSize
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=pageSize')]
    public ?int $pageSize = null;
    
    /**
     * The maximum number of results to return per page.
     * 
     * Deprecated, please use `pageSize` instead.
     * 
     * 
     * @var ?int $pageSizeDeprecated
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=page_size')]
    public ?int $pageSizeDeprecated = null;
    
    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     * 
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     * Deprecated, please use `cursor` instead.
     * 
     * 
     * @var ?string $paginationToken
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=pagination_token')]
    public ?string $paginationToken = null;
    
    /**
     * Find transactions by reference field.
     * 
     * @var ?string $reference
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=reference')]
    public ?string $reference = null;
    
    /**
     * Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
     * 
     * @var ?string $source
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=source')]
    public ?string $source = null;
    
    /**
     * Filter transactions that occurred after this timestamp.
     * 
     * The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
     * 
     * 
     * @var ?\DateTime $startTime
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=startTime,dateTimeFormat=Y-m-d\TH:i:s.up')]
    public ?\DateTime $startTime = null;
    
    /**
     * Filter transactions that occurred after this timestamp.
     * 
     * The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
     * Deprecated, please use `startTime` instead.
     * 
     * 
     * @var ?\DateTime $startTimeDeprecated
     */
	#[SpeakeasyMetadata('queryParam:style=form,explode=true,name=start_time,dateTimeFormat=Y-m-d\TH:i:s.up')]
    public ?\DateTime $startTimeDeprecated = null;
    
	public function __construct()
	{
		$this->account = null;
		$this->after = null;
		$this->cursor = null;
		$this->destination = null;
		$this->endTime = null;
		$this->endTimeDeprecated = null;
		$this->ledger = "";
		$this->metadata = null;
		$this->pageSize = null;
		$this->pageSizeDeprecated = null;
		$this->paginationToken = null;
		$this->reference = null;
		$this->source = null;
		$this->startTime = null;
		$this->startTimeDeprecated = null;
	}
}
