# UpdateAppClient

The AWS AppFabric for productivity feature is in preview and is subject to change.

Updates an AppClient.

###### Topics

- [Request body](#API_UpdateAppClient_request)

- [Response elements](#API_UpdateAppClient_response)

## Request body

The request accepts the following data in JSON format.

ParameterDescription

**appClientIdentifier**

The Amazon Resource Name (ARN) or Universal Unique
Identifier (UUID) of the AppClient to use for the
request.

Length Constraints: Minimum length of 1. Maximum length of
1011.

Pattern:
`arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**redirectUrls**

The URI to redirect end users to after authorization. You
can add up to 5 redirectUrls. For example,
`https://localhost:8080`.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of
5 items.

Length Constraints: Minimum length of 1. Maximum length of
2048.

Pattern:
`(http|https):\/\/[-a-zA-Z0-9_:.\/]+`

## Response elements

If the action is successful, the service sends back an HTTP 200
response.

The following data is returned in JSON format by the service.

ParameterDescription

**appClient**

Contains information about an AppClient.

Type: [AppClient](api-appclient.md)
object

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Token

Data types

All content copied from https://docs.aws.amazon.com/.
