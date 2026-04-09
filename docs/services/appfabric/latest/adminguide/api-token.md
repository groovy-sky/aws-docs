# Token

The AWS AppFabric for productivity feature is in preview and is subject to change.

Contains information that allows AppClients to exchange an authorization code for
an access token.

###### Topics

- [Request body](#API_Token_request)

- [Response elements](#API_Token_response)

## Request body

The request accepts the following data in JSON format.

ParameterDescription

**code**

The authorization code received from the authorization
endpoint.

Type: String

Length Constraints: Minimum length of 1. Maximum length of
2048.

Required: No

**grant\_type**

The grant type for the token. Must be either
`authorization_code` or
`refresh_token`.

Type: String

Required: Yes

**app\_client\_id**

The ID of the AppClient.

Type: String

Pattern:
`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**redirect\_uri**

The redirect URI passed to the authorization
endpoint.

Type: String

Required: No

**refresh\_token**

The refresh token received from the initial token
request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of
4096.

Required: No

## Response elements

If the action is successful, the service sends back an HTTP 200
response.

The following data is returned in JSON format by the service.

ParameterDescription

**appfabric\_user\_id**

The ID of the user for the token. This is returned only
for requests that use the `authorization_code`
grant type.

Type: String

**expires\_in**

The number of seconds until the token expires.

Type: Long

**refresh\_token**

The refresh token to use for a subsequent request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of
2048.

**token**

The access token.

Type: String

Length Constraints: Minimum length of 1. Maximum length of
2048.

**token\_type**

The token type.

Type: String

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutFeedback

UpdateAppClient

All content copied from https://docs.aws.amazon.com/.
