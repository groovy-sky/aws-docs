# DescribeContributorInsights

Returns information about contributor insights for a given table or global secondary
index.

## Request Syntax

```nohighlight

{
   "IndexName": "string",
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[TableName](#API_DescribeContributorInsights_RequestSyntax)**

The name of the table to describe. You can also provide the Amazon Resource Name (ARN) of the table in
this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[IndexName](#API_DescribeContributorInsights_RequestSyntax)**

The name of the global secondary index to describe, if applicable.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## Response Syntax

```nohighlight

{
   "ContributorInsightsMode": "string",
   "ContributorInsightsRuleList": [ "string" ],
   "ContributorInsightsStatus": "string",
   "FailureException": {
      "ExceptionDescription": "string",
      "ExceptionName": "string"
   },
   "IndexName": "string",
   "LastUpdateDateTime": number,
   "TableName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ContributorInsightsMode](#API_DescribeContributorInsights_ResponseSyntax)**

The mode of CloudWatch Contributor Insights for DynamoDB that determines
which events are emitted. Can be set to track all access and throttled events or throttled
events only.

Type: String

Valid Values: `ACCESSED_AND_THROTTLED_KEYS | THROTTLED_KEYS`

**[ContributorInsightsRuleList](#API_DescribeContributorInsights_ResponseSyntax)**

List of names of the associated contributor insights rules.

Type: Array of strings

Pattern: `[A-Za-z0-9][A-Za-z0-9\-\_\.]{0,126}[A-Za-z0-9]`

**[ContributorInsightsStatus](#API_DescribeContributorInsights_ResponseSyntax)**

Current status of contributor insights.

Type: String

Valid Values: `ENABLING | ENABLED | DISABLING | DISABLED | FAILED`

**[FailureException](#API_DescribeContributorInsights_ResponseSyntax)**

Returns information about the last failure that was encountered.

The most common exceptions for a FAILED status are:

- LimitExceededException - Per-account Amazon CloudWatch Contributor Insights
rule limit reached. Please disable Contributor Insights for other tables/indexes
OR disable Contributor Insights rules before retrying.

- AccessDeniedException - Amazon CloudWatch Contributor Insights rules cannot be
modified due to insufficient permissions.

- AccessDeniedException - Failed to create service-linked role for Contributor
Insights due to insufficient permissions.

- InternalServerError - Failed to create Amazon CloudWatch Contributor Insights
rules. Please retry request.

Type: [FailureException](api-failureexception.md) object

**[IndexName](#API_DescribeContributorInsights_ResponseSyntax)**

The name of the global secondary index being described.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

**[LastUpdateDateTime](#API_DescribeContributorInsights_ResponseSyntax)**

Timestamp of the last time the status was changed.

Type: Timestamp

**[TableName](#API_DescribeContributorInsights_ResponseSyntax)**

The name of the table being described.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/describecontributorinsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/describecontributorinsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeContinuousBackups

DescribeEndpoints

All content copied from https://docs.aws.amazon.com/.
