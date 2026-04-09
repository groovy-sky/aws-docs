# ListFlows

Lists all of the flows associated with your account.

## Request Syntax

```nohighlight

POST /list-flows HTTP/1.1
Content-type: application/json

{
   "maxResults": number,
   "nextToken": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[maxResults](#API_ListFlows_RequestSyntax)**

Specifies the maximum number of items that should be returned in the result set.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[nextToken](#API_ListFlows_RequestSyntax)**

The pagination token for next page of data.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "flows": [
      {
         "createdAt": number,
         "createdBy": "string",
         "description": "string",
         "destinationConnectorLabel": "string",
         "destinationConnectorType": "string",
         "flowArn": "string",
         "flowName": "string",
         "flowStatus": "string",
         "lastRunExecutionDetails": {
            "mostRecentExecutionMessage": "string",
            "mostRecentExecutionStatus": "string",
            "mostRecentExecutionTime": number
         },
         "lastUpdatedAt": number,
         "lastUpdatedBy": "string",
         "sourceConnectorLabel": "string",
         "sourceConnectorType": "string",
         "tags": {
            "string" : "string"
         },
         "triggerType": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[flows](#API_ListFlows_ResponseSyntax)**

The list of flows associated with your account.

Type: Array of [FlowDefinition](api-flowdefinition.md) objects

**[nextToken](#API_ListFlows_ResponseSyntax)**

The pagination token for next page of data.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `\S+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ValidationException**

The request has invalid or missing parameters.

HTTP Status Code: 400

## Examples

### ListFlows examples

This example shows a sample request for the `ListFlows` API. In the second
sample, note that `MaxResults` will show a number between 1 and 100.

#### Sample Request

```json

{
  "flowList": [
    {
      "createdAt": "created_time_value",
      "createdBy": "arn:aws:iam::<AccountId>:user/BetaTestUser",
      "description": "test flow 1 description",
      "destinationConnectorType": "S3",
      "destinationConnectorLabel": "MyCustomDestinationConnector",
      "flowArn": "arn:aws:appflow:region:<AccountId>:flow/test-flow-1",
      "flowName": "test-flow-1",
      "flowStatus": "Active",
      "lastRunExecutionDetails":
      {
        "mostRecentExecutionMessage": "Successfully ran the flow",
        "mostRecentExecutionStatus": "Successful",
        "mostRecentExecutionTime": "execution_time_value"
      },
      "lastUpdatedAt": "lastupdated_time_value",
      "lastUpdatedBy": "arn:aws:iam::<AccountId>:user/BetaTestUser",
      "sourceConnectorType": "Salesforce",
      "sourceConnectorLabel": "MyCustomSourceConnector",
      "tags":
      {
        "internalId": "<InternalId>",
        "resourceArn": "arn:aws:appflow:region:<AccountId>:flow/test-flow-1"
      },
      "triggerType": "OnDemand"
    }
  ],
  "nextToken": "next_token_value"
}
```

```json

{
  "maxResults": 1
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appflow-2020-08-23/listflows.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appflow-2020-08-23/listflows.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/listflows.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appflow-2020-08-23/listflows.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/listflows.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appflow-2020-08-23/listflows.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appflow-2020-08-23/listflows.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appflow-2020-08-23/listflows.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appflow-2020-08-23/listflows.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/listflows.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListConnectors

ListTagsForResource

All content copied from https://docs.aws.amazon.com/.
