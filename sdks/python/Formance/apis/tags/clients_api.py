# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: v1.0.20230228
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from Formance.paths.api_auth_clients_client_id_scopes_scope_id.put import AddScopeToClient
from Formance.paths.api_auth_clients.post import CreateClient
from Formance.paths.api_auth_clients_client_id_secrets.post import CreateSecret
from Formance.paths.api_auth_clients_client_id.delete import DeleteClient
from Formance.paths.api_auth_clients_client_id_scopes_scope_id.delete import DeleteScopeFromClient
from Formance.paths.api_auth_clients_client_id_secrets_secret_id.delete import DeleteSecret
from Formance.paths.api_auth_clients.get import ListClients
from Formance.paths.api_auth_clients_client_id.get import ReadClient
from Formance.paths.api_auth_clients_client_id.put import UpdateClient


class ClientsApi(
    AddScopeToClient,
    CreateClient,
    CreateSecret,
    DeleteClient,
    DeleteScopeFromClient,
    DeleteSecret,
    ListClients,
    ReadClient,
    UpdateClient,
):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """
    pass
