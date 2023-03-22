"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import errorresponse as shared_errorresponse
from ..shared import transactionscursorresponse as shared_transactionscursorresponse
from datetime import datetime
from typing import Any, Optional


@dataclasses.dataclass
class ListTransactionsRequest:
    
    ledger: str = dataclasses.field(metadata={'path_param': { 'field_name': 'ledger', 'style': 'simple', 'explode': False }})
    r"""Name of the ledger."""  
    account: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'account', 'style': 'form', 'explode': True }})
    r"""Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $)."""  
    after: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'after', 'style': 'form', 'explode': True }})
    r"""Pagination cursor, will return transactions after given txid (in descending order)."""  
    cursor: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'cursor', 'style': 'form', 'explode': True }})
    r"""Parameter used in pagination requests. Maximum page size is set to 15.
    Set to the value of next for the next page of results.
    Set to the value of previous for the previous page of results.
    No other parameters can be set when this parameter is set.
    
    """  
    destination: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'destination', 'style': 'form', 'explode': True }})
    r"""Filter transactions with postings involving given account at destination (regular expression placed between ^ and $)."""  
    end_time: Optional[datetime] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'endTime', 'style': 'form', 'explode': True }})
    r"""Filter transactions that occurred before this timestamp.
    The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
    
    """  
    end_time_deprecated: Optional[datetime] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'end_time', 'style': 'form', 'explode': True }})
    r"""Filter transactions that occurred before this timestamp.
    The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
    Deprecated, please use `endTime` instead.
    
    """  
    metadata: Optional[dict[str, Any]] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'metadata', 'style': 'deepObject', 'explode': True }})
    r"""Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below."""  
    page_size: Optional[int] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'pageSize', 'style': 'form', 'explode': True }})
    r"""The maximum number of results to return per page.
    
    """  
    page_size_deprecated: Optional[int] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'page_size', 'style': 'form', 'explode': True }})
    r"""The maximum number of results to return per page.
    Deprecated, please use `pageSize` instead.
    
    """  
    pagination_token: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'pagination_token', 'style': 'form', 'explode': True }})
    r"""Parameter used in pagination requests. Maximum page size is set to 15.
    Set to the value of next for the next page of results.
    Set to the value of previous for the previous page of results.
    No other parameters can be set when this parameter is set.
    Deprecated, please use `cursor` instead.
    
    """  
    reference: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'reference', 'style': 'form', 'explode': True }})
    r"""Find transactions by reference field."""  
    source: Optional[str] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'source', 'style': 'form', 'explode': True }})
    r"""Filter transactions with postings involving given account at source (regular expression placed between ^ and $)."""  
    start_time: Optional[datetime] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'startTime', 'style': 'form', 'explode': True }})
    r"""Filter transactions that occurred after this timestamp.
    The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
    
    """  
    start_time_deprecated: Optional[datetime] = dataclasses.field(default=None, metadata={'query_param': { 'field_name': 'start_time', 'style': 'form', 'explode': True }})
    r"""Filter transactions that occurred after this timestamp.
    The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
    Deprecated, please use `startTime` instead.
    
    """  
    

@dataclasses.dataclass
class ListTransactionsResponse:
    
    content_type: str = dataclasses.field()  
    status_code: int = dataclasses.field()  
    error_response: Optional[shared_errorresponse.ErrorResponse] = dataclasses.field(default=None)
    r"""Error"""  
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)  
    transactions_cursor_response: Optional[shared_transactionscursorresponse.TransactionsCursorResponse] = dataclasses.field(default=None)
    r"""OK"""  
    