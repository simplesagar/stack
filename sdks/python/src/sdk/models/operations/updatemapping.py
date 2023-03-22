"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import errorresponse as shared_errorresponse
from ..shared import mapping as shared_mapping
from ..shared import mappingresponse as shared_mappingresponse
from typing import Optional


@dataclasses.dataclass
class UpdateMappingRequest:
    
    ledger: str = dataclasses.field(metadata={'path_param': { 'field_name': 'ledger', 'style': 'simple', 'explode': False }})
    r"""Name of the ledger."""  
    mapping: shared_mapping.Mapping = dataclasses.field(metadata={'request': { 'media_type': 'application/json' }})  
    

@dataclasses.dataclass
class UpdateMappingResponse:
    
    content_type: str = dataclasses.field()  
    status_code: int = dataclasses.field()  
    error_response: Optional[shared_errorresponse.ErrorResponse] = dataclasses.field(default=None)
    r"""Error"""  
    mapping_response: Optional[shared_mappingresponse.MappingResponse] = dataclasses.field(default=None)
    r"""OK"""  
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)  
    