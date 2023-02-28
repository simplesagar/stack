/**
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * OpenAPI spec version: v1.0.20230228
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class Query {
    'ledgers'?: Array<string>;
    'after'?: Array<string>;
    'pageSize'?: number;
    'terms'?: Array<string>;
    'sort'?: string;
    'policy'?: string;
    'target'?: string;
    'cursor'?: string;
    'raw'?: any;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "ledgers",
            "baseName": "ledgers",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "after",
            "baseName": "after",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "pageSize",
            "baseName": "pageSize",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "terms",
            "baseName": "terms",
            "type": "Array<string>",
            "format": ""
        },
        {
            "name": "sort",
            "baseName": "sort",
            "type": "string",
            "format": ""
        },
        {
            "name": "policy",
            "baseName": "policy",
            "type": "string",
            "format": ""
        },
        {
            "name": "target",
            "baseName": "target",
            "type": "string",
            "format": ""
        },
        {
            "name": "cursor",
            "baseName": "cursor",
            "type": "string",
            "format": ""
        },
        {
            "name": "raw",
            "baseName": "raw",
            "type": "any",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Query.attributeTypeMap;
    }

    public constructor() {
    }
}

