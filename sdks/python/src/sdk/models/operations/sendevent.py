"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import error as shared_error
from dataclasses_json import Undefined, dataclass_json
from sdk import utils
from typing import Optional


@dataclass_json(undefined=Undefined.EXCLUDE)
@dataclasses.dataclass
class SendEventRequestBody:
    
    name: str = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('name') }})  
    

@dataclasses.dataclass
class SendEventRequest:
    
    instance_id: str = dataclasses.field(metadata={'path_param': { 'field_name': 'instanceID', 'style': 'simple', 'explode': False }})
    r"""The instance id"""  
    request_body: Optional[SendEventRequestBody] = dataclasses.field(default=None, metadata={'request': { 'media_type': 'application/json' }})  
    

@dataclasses.dataclass
class SendEventResponse:
    
    content_type: str = dataclasses.field()  
    status_code: int = dataclasses.field()  
    error: Optional[shared_error.Error] = dataclasses.field(default=None)
    r"""General error"""  
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)  
    