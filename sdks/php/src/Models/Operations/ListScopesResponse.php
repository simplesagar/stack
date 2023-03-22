<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class ListScopesResponse
{
	
    public string $contentType;
    
    /**
     * List of scopes
     * 
     * @var ?\formance\stack\Models\Shared\ListScopesResponse $listScopesResponse
     */
	
    public ?\formance\stack\Models\Shared\ListScopesResponse $listScopesResponse = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->listScopesResponse = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}
