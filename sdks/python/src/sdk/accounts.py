"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

import requests as requests_http
from . import utils
from sdk.models import operations, shared
from typing import Optional

class Accounts:
    _client: requests_http.Session
    _security_client: requests_http.Session
    _server_url: str
    _language: str
    _sdk_version: str
    _gen_version: str

    def __init__(self, client: requests_http.Session, security_client: requests_http.Session, server_url: str, language: str, sdk_version: str, gen_version: str) -> None:
        self._client = client
        self._security_client = security_client
        self._server_url = server_url
        self._language = language
        self._sdk_version = sdk_version
        self._gen_version = gen_version
        
    def add_metadata_to_account(self, request: operations.AddMetadataToAccountRequest) -> operations.AddMetadataToAccountResponse:
        r"""Add metadata to an account"""
        base_url = self._server_url
        
        url = utils.generate_url(operations.AddMetadataToAccountRequest, base_url, '/api/ledger/{ledger}/accounts/{address}/metadata', request)
        
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request_body", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        
        client = self._security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.AddMetadataToAccountResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 204:
            pass
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    def count_accounts(self, request: operations.CountAccountsRequest) -> operations.CountAccountsResponse:
        r"""Count the accounts from a ledger"""
        base_url = self._server_url
        
        url = utils.generate_url(operations.CountAccountsRequest, base_url, '/api/ledger/{ledger}/accounts', request)
        
        query_params = utils.get_query_params(operations.CountAccountsRequest, request)
        
        client = self._security_client
        
        http_res = client.request('HEAD', url, params=query_params)
        content_type = http_res.headers.get('Content-Type')

        res = operations.CountAccountsResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            res.headers = http_res.headers
            
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    def get_account(self, request: operations.GetAccountRequest) -> operations.GetAccountResponse:
        r"""Get account by its address"""
        base_url = self._server_url
        
        url = utils.generate_url(operations.GetAccountRequest, base_url, '/api/ledger/{ledger}/accounts/{address}', request)
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetAccountResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.AccountResponse])
                res.account_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    def list_accounts(self, request: operations.ListAccountsRequest) -> operations.ListAccountsResponse:
        r"""List accounts from a ledger
        List accounts from a ledger, sorted by address in descending order.
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.ListAccountsRequest, base_url, '/api/ledger/{ledger}/accounts', request)
        
        query_params = utils.get_query_params(operations.ListAccountsRequest, request)
        
        client = self._security_client
        
        http_res = client.request('GET', url, params=query_params)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ListAccountsResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.AccountsCursorResponse])
                res.accounts_cursor_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    