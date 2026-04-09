# GetAppClient

The AWS AppFabric for productivity feature is in preview and is subject to change.

Returns information about an AppClient.

###### Topics

- [Request body](#API_GetAppClient_request)

- [Response elements](#API_GetAppClient_response)

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

DeleteAppClient

ListActionableInsights

All content copied from https://docs.aws.amazon.com/.
