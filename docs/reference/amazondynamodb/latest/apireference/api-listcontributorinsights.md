# ListContributorInsights

Returns a list of ContributorInsightsSummary for a table and all its global secondary
indexes.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[MaxResults](#API_ListContributorInsights_RequestSyntax)**

Maximum number of results to return per page.

Type: Integer

Valid Range: Maximum value of 100.

Required: No

**[NextToken](#API_ListContributorInsights_RequestSyntax)**

A token to for the desired page, if there is one.

Type: String

Required: No

**[TableName](#API_ListContributorInsights_RequestSyntax)**

The name of the table. You can also provide the Amazon Resource Name (ARN) of the table in this
parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## Response Syntax

```nohighlight

{
   "ContributorInsightsSummaries": [
      {
         "ContributorInsightsMode": "string",
         "ContributorInsightsStatus": "string",
         "IndexName": "string",
         "TableName": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ContributorInsightsSummaries](#API_ListContributorInsights_ResponseSyntax)**

A list of ContributorInsightsSummary.

Type: Array of [ContributorInsightsSummary](api-contributorinsightssummary.md) objects

**[NextToken](#API_ListContributorInsights_ResponseSyntax)**

A token to go to the next page if there is one.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/listcontributorinsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/listcontributorinsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListBackups

ListExports

All content copied from https://docs.aws.amazon.com/.
