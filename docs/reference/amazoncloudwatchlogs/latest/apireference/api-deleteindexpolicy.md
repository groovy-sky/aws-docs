# DeleteIndexPolicy

Deletes a log-group level field index policy that was applied to a single log group. The
indexing of the log events that happened before you delete the policy will still be used for
as many as 30 days to improve CloudWatch Logs Insights queries.

If the deleted policy included facet configurations, those facets will no longer be
available for interactive exploration in the CloudWatch Logs Insights console for this log
group. However, facet data is retained for up to 30 days.

You can't use this operation to delete an account-level index policy. Instead, use [DeleteAccountPolicy](api-deleteaccountpolicy.md).

If you delete a log-group level field index policy and there is an account-level field
index policy, in a few minutes the log group begins using that account-wide policy to index
new incoming log events. This operation only affects log group-level policies, including any
facet configurations, and preserves any data source-based account policies that may apply to
the log group.

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_DeleteIndexPolicy_RequestSyntax)**

The log group to delete the index policy for. You can specify either the name or the ARN
of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deleteindexpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deleteindexpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteDestination

DeleteIntegration
