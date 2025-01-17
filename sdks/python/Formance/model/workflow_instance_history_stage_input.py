# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: develop
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from Formance import schemas  # noqa: F401


class WorkflowInstanceHistoryStageInput(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
        
            @staticmethod
            def GetAccount() -> typing.Type['ActivityGetAccount']:
                return ActivityGetAccount
        
            @staticmethod
            def CreateTransaction() -> typing.Type['ActivityCreateTransaction']:
                return ActivityCreateTransaction
        
            @staticmethod
            def RevertTransaction() -> typing.Type['ActivityRevertTransaction']:
                return ActivityRevertTransaction
        
            @staticmethod
            def StripeTransfer() -> typing.Type['StripeTransferRequest']:
                return StripeTransferRequest
        
            @staticmethod
            def GetPayment() -> typing.Type['ActivityGetPayment']:
                return ActivityGetPayment
        
            @staticmethod
            def ConfirmHold() -> typing.Type['ActivityConfirmHold']:
                return ActivityConfirmHold
        
            @staticmethod
            def CreditWallet() -> typing.Type['ActivityCreditWallet']:
                return ActivityCreditWallet
        
            @staticmethod
            def DebitWallet() -> typing.Type['ActivityDebitWallet']:
                return ActivityDebitWallet
        
            @staticmethod
            def GetWallet() -> typing.Type['ActivityGetWallet']:
                return ActivityGetWallet
        
            @staticmethod
            def VoidHold() -> typing.Type['ActivityVoidHold']:
                return ActivityVoidHold
            __annotations__ = {
                "GetAccount": GetAccount,
                "CreateTransaction": CreateTransaction,
                "RevertTransaction": RevertTransaction,
                "StripeTransfer": StripeTransfer,
                "GetPayment": GetPayment,
                "ConfirmHold": ConfirmHold,
                "CreditWallet": CreditWallet,
                "DebitWallet": DebitWallet,
                "GetWallet": GetWallet,
                "VoidHold": VoidHold,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["GetAccount"]) -> 'ActivityGetAccount': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["CreateTransaction"]) -> 'ActivityCreateTransaction': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["RevertTransaction"]) -> 'ActivityRevertTransaction': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["StripeTransfer"]) -> 'StripeTransferRequest': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["GetPayment"]) -> 'ActivityGetPayment': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["ConfirmHold"]) -> 'ActivityConfirmHold': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["CreditWallet"]) -> 'ActivityCreditWallet': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["DebitWallet"]) -> 'ActivityDebitWallet': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["GetWallet"]) -> 'ActivityGetWallet': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["VoidHold"]) -> 'ActivityVoidHold': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["GetAccount", "CreateTransaction", "RevertTransaction", "StripeTransfer", "GetPayment", "ConfirmHold", "CreditWallet", "DebitWallet", "GetWallet", "VoidHold", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["GetAccount"]) -> typing.Union['ActivityGetAccount', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["CreateTransaction"]) -> typing.Union['ActivityCreateTransaction', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["RevertTransaction"]) -> typing.Union['ActivityRevertTransaction', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["StripeTransfer"]) -> typing.Union['StripeTransferRequest', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["GetPayment"]) -> typing.Union['ActivityGetPayment', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["ConfirmHold"]) -> typing.Union['ActivityConfirmHold', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["CreditWallet"]) -> typing.Union['ActivityCreditWallet', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["DebitWallet"]) -> typing.Union['ActivityDebitWallet', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["GetWallet"]) -> typing.Union['ActivityGetWallet', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["VoidHold"]) -> typing.Union['ActivityVoidHold', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["GetAccount", "CreateTransaction", "RevertTransaction", "StripeTransfer", "GetPayment", "ConfirmHold", "CreditWallet", "DebitWallet", "GetWallet", "VoidHold", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        GetAccount: typing.Union['ActivityGetAccount', schemas.Unset] = schemas.unset,
        CreateTransaction: typing.Union['ActivityCreateTransaction', schemas.Unset] = schemas.unset,
        RevertTransaction: typing.Union['ActivityRevertTransaction', schemas.Unset] = schemas.unset,
        StripeTransfer: typing.Union['StripeTransferRequest', schemas.Unset] = schemas.unset,
        GetPayment: typing.Union['ActivityGetPayment', schemas.Unset] = schemas.unset,
        ConfirmHold: typing.Union['ActivityConfirmHold', schemas.Unset] = schemas.unset,
        CreditWallet: typing.Union['ActivityCreditWallet', schemas.Unset] = schemas.unset,
        DebitWallet: typing.Union['ActivityDebitWallet', schemas.Unset] = schemas.unset,
        GetWallet: typing.Union['ActivityGetWallet', schemas.Unset] = schemas.unset,
        VoidHold: typing.Union['ActivityVoidHold', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'WorkflowInstanceHistoryStageInput':
        return super().__new__(
            cls,
            *_args,
            GetAccount=GetAccount,
            CreateTransaction=CreateTransaction,
            RevertTransaction=RevertTransaction,
            StripeTransfer=StripeTransfer,
            GetPayment=GetPayment,
            ConfirmHold=ConfirmHold,
            CreditWallet=CreditWallet,
            DebitWallet=DebitWallet,
            GetWallet=GetWallet,
            VoidHold=VoidHold,
            _configuration=_configuration,
            **kwargs,
        )

from Formance.model.activity_confirm_hold import ActivityConfirmHold
from Formance.model.activity_create_transaction import ActivityCreateTransaction
from Formance.model.activity_credit_wallet import ActivityCreditWallet
from Formance.model.activity_debit_wallet import ActivityDebitWallet
from Formance.model.activity_get_account import ActivityGetAccount
from Formance.model.activity_get_payment import ActivityGetPayment
from Formance.model.activity_get_wallet import ActivityGetWallet
from Formance.model.activity_revert_transaction import ActivityRevertTransaction
from Formance.model.activity_void_hold import ActivityVoidHold
from Formance.model.stripe_transfer_request import StripeTransferRequest
