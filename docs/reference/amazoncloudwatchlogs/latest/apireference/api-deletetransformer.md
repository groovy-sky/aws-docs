# DeleteTransformer

Deletes the log transformer for the specified log group. As soon as you do this, the
transformation of incoming log events according to that transformer stops. If this account has
an account-level transformer that applies to this log group, the log group begins using that
account-level transformer when this log-group level transformer is deleted.

After you delete a transformer, be sure to edit any metric filters or subscription filters
that relied on the transformed versions of the log events.

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_DeleteTransformer_RequestSyntax)**

Specify either the name or ARN of the log group to delete the transformer for. If the log
group is in a source account and you are using a monitoring account, you must use the log
group ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperationException**

The operation is not valid on the specified resource.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deletetransformer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deletetransformer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteSubscriptionFilter

DescribeAccountPolicies
