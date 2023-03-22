/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.formance.formance_sdk.utils.HTTPClient;
import com.formance.formance_sdk.utils.HTTPRequest;
import com.formance.formance_sdk.utils.JSON;
import com.formance.formance_sdk.utils.SerializedBody;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;

public class Auth {
	
	private HTTPClient _defaultClient;
	private HTTPClient _securityClient;
	private String _serverUrl;
	private String _language;
	private String _sdkVersion;
	private String _genVersion;

	public Auth(HTTPClient defaultClient, HTTPClient securityClient, String serverUrl, String language, String sdkVersion, String genVersion) {
		this._defaultClient = defaultClient;
		this._securityClient = securityClient;
		this._serverUrl = serverUrl;
		this._language = language;
		this._sdkVersion = sdkVersion;
		this._genVersion = genVersion;
	}

    /**
     * Add scope to client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.AddScopeToClientResponse addScopeToClient(com.formance.formance_sdk.models.operations.AddScopeToClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.AddScopeToClientRequest.class, baseUrl, "/api/auth/clients/{clientId}/scopes/{scopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("PUT");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.AddScopeToClientResponse res = new com.formance.formance_sdk.models.operations.AddScopeToClientResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Add a transient scope to a scope
     * Add a transient scope to a scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.AddTransientScopeResponse addTransientScope(com.formance.formance_sdk.models.operations.AddTransientScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.AddTransientScopeRequest.class, baseUrl, "/api/auth/scopes/{scopeId}/transient/{transientScopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("PUT");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.AddTransientScopeResponse res = new com.formance.formance_sdk.models.operations.AddTransientScopeResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Create client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreateClientResponse createClient(com.formance.formance_sdk.models.shared.CreateClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/auth/clients");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "request", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreateClientResponse res = new com.formance.formance_sdk.models.operations.CreateClientResponse() {{
            createClientResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 201) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.CreateClientResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.CreateClientResponse.class);
                res.createClientResponse = out;
            }
        }

        return res;
    }

    /**
     * Create scope
     * Create scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreateScopeResponse createScope(com.formance.formance_sdk.models.shared.CreateScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/auth/scopes");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "request", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreateScopeResponse res = new com.formance.formance_sdk.models.operations.CreateScopeResponse() {{
            createScopeResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 201) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.CreateScopeResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.CreateScopeResponse.class);
                res.createScopeResponse = out;
            }
        }

        return res;
    }

    /**
     * Add a secret to a client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreateSecretResponse createSecret(com.formance.formance_sdk.models.operations.CreateSecretRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.CreateSecretRequest.class, baseUrl, "/api/auth/clients/{clientId}/secrets", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "createSecretRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreateSecretResponse res = new com.formance.formance_sdk.models.operations.CreateSecretResponse() {{
            createSecretResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.CreateSecretResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.CreateSecretResponse.class);
                res.createSecretResponse = out;
            }
        }

        return res;
    }

    /**
     * Delete client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DeleteClientResponse deleteClient(com.formance.formance_sdk.models.operations.DeleteClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DeleteClientRequest.class, baseUrl, "/api/auth/clients/{clientId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("DELETE");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DeleteClientResponse res = new com.formance.formance_sdk.models.operations.DeleteClientResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Delete scope
     * Delete scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DeleteScopeResponse deleteScope(com.formance.formance_sdk.models.operations.DeleteScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DeleteScopeRequest.class, baseUrl, "/api/auth/scopes/{scopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("DELETE");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DeleteScopeResponse res = new com.formance.formance_sdk.models.operations.DeleteScopeResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Delete scope from client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DeleteScopeFromClientResponse deleteScopeFromClient(com.formance.formance_sdk.models.operations.DeleteScopeFromClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DeleteScopeFromClientRequest.class, baseUrl, "/api/auth/clients/{clientId}/scopes/{scopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("DELETE");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DeleteScopeFromClientResponse res = new com.formance.formance_sdk.models.operations.DeleteScopeFromClientResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Delete a secret from a client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DeleteSecretResponse deleteSecret(com.formance.formance_sdk.models.operations.DeleteSecretRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DeleteSecretRequest.class, baseUrl, "/api/auth/clients/{clientId}/secrets/{secretId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("DELETE");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DeleteSecretResponse res = new com.formance.formance_sdk.models.operations.DeleteSecretResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * Delete a transient scope from a scope
     * Delete a transient scope from a scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DeleteTransientScopeResponse deleteTransientScope(com.formance.formance_sdk.models.operations.DeleteTransientScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DeleteTransientScopeRequest.class, baseUrl, "/api/auth/scopes/{scopeId}/transient/{transientScopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("DELETE");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DeleteTransientScopeResponse res = new com.formance.formance_sdk.models.operations.DeleteTransientScopeResponse() {{
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }

        return res;
    }

    /**
     * List clients
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ListClientsResponse listClients() throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/auth/clients");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ListClientsResponse res = new com.formance.formance_sdk.models.operations.ListClientsResponse() {{
            listClientsResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ListClientsResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ListClientsResponse.class);
                res.listClientsResponse = out;
            }
        }

        return res;
    }

    /**
     * List scopes
     * List Scopes
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ListScopesResponse listScopes() throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/auth/scopes");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ListScopesResponse res = new com.formance.formance_sdk.models.operations.ListScopesResponse() {{
            listScopesResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ListScopesResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ListScopesResponse.class);
                res.listScopesResponse = out;
            }
        }

        return res;
    }

    /**
     * List users
     * List users
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ListUsersResponse listUsers() throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/auth/users");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ListUsersResponse res = new com.formance.formance_sdk.models.operations.ListUsersResponse() {{
            listUsersResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ListUsersResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ListUsersResponse.class);
                res.listUsersResponse = out;
            }
        }

        return res;
    }

    /**
     * Read client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ReadClientResponse readClient(com.formance.formance_sdk.models.operations.ReadClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.ReadClientRequest.class, baseUrl, "/api/auth/clients/{clientId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ReadClientResponse res = new com.formance.formance_sdk.models.operations.ReadClientResponse() {{
            readClientResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ReadClientResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ReadClientResponse.class);
                res.readClientResponse = out;
            }
        }

        return res;
    }

    /**
     * Read scope
     * Read scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ReadScopeResponse readScope(com.formance.formance_sdk.models.operations.ReadScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.ReadScopeRequest.class, baseUrl, "/api/auth/scopes/{scopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ReadScopeResponse res = new com.formance.formance_sdk.models.operations.ReadScopeResponse() {{
            readScopeResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ReadScopeResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ReadScopeResponse.class);
                res.readScopeResponse = out;
            }
        }

        return res;
    }

    /**
     * Read user
     * Read user
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ReadUserResponse readUser(com.formance.formance_sdk.models.operations.ReadUserRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.ReadUserRequest.class, baseUrl, "/api/auth/users/{userId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ReadUserResponse res = new com.formance.formance_sdk.models.operations.ReadUserResponse() {{
            readUserResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ReadUserResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ReadUserResponse.class);
                res.readUserResponse = out;
            }
        }

        return res;
    }

    /**
     * Update client
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.UpdateClientResponse updateClient(com.formance.formance_sdk.models.operations.UpdateClientRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.UpdateClientRequest.class, baseUrl, "/api/auth/clients/{clientId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("PUT");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "updateClientRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.UpdateClientResponse res = new com.formance.formance_sdk.models.operations.UpdateClientResponse() {{
            updateClientResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.UpdateClientResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.UpdateClientResponse.class);
                res.updateClientResponse = out;
            }
        }

        return res;
    }

    /**
     * Update scope
     * Update scope
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.UpdateScopeResponse updateScope(com.formance.formance_sdk.models.operations.UpdateScopeRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.UpdateScopeRequest.class, baseUrl, "/api/auth/scopes/{scopeId}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("PUT");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "updateScopeRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.UpdateScopeResponse res = new com.formance.formance_sdk.models.operations.UpdateScopeResponse() {{
            updateScopeResponse = null;
        }};
        res.statusCode = httpRes.statusCode();
        res.contentType = contentType;
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.UpdateScopeResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.UpdateScopeResponse.class);
                res.updateScopeResponse = out;
            }
        }

        return res;
    }
}