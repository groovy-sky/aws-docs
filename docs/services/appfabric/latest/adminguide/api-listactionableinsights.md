# ListActionableInsights

The AWS AppFabric for productivity feature is in preview and is subject to change.

Lists the most important actionable email messages, tasks, and other
updates.

###### Topics

- [Request body](#API_ListActionableInsights_request)

- [Response elements](#API_ListActionableInsights_response)

## Request body

The request accepts the following data in JSON format.

ParameterDescription

**nextToken**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

## Response elements

If the action is successful, the service sends back an HTTP 201
response.

The following data is returned in JSON format by the service.

ParameterDescription

**ActionableInsightsList**

Lists the actionable insights, including a title,
description, actions, and created timestamp. For more
information, see [ActionableInsights](api-actionableinsights.md).

**nextToken**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

Type: String

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAppClient

ListAppClients

All content copied from https://docs.aws.amazon.com/.
