# DescribeIndexPolicies

Returns the field index policies of the specified log group. For more information about
field index policies, see [PutIndexPolicy](api-putindexpolicy.md).

If a specified log group has a log-group level index policy, that policy is returned by
this operation.

If a specified log group doesn't have a log-group level index policy, but an account-wide
index policy applies to it, that account-wide policy is returned by this operation.

To find information about only account-level policies, use [DescribeAccountPolicies](api-describeaccountpolicies.md) instead.

## Request Syntax

```nohighlight

{
   "logGroupIdentifiers": [ "string" ],
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifiers](#API_DescribeIndexPolicies_RequestSyntax)**

An array containing the name or ARN of the log group that you want to retrieve field index
policies for.

Type: Array of strings

Array Members: Fixed number of 1 item.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[nextToken](#API_DescribeIndexPolicies_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "indexPolicies": [
      {
         "lastUpdateTime": number,
         "logGroupIdentifier": "string",
         "policyDocument": "string",
         "policyName": "string",
         "source": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[indexPolicies](#API_DescribeIndexPolicies_ResponseSyntax)**

An array containing the field index policies.

Type: Array of [IndexPolicy](api-indexpolicy.md) objects

**[nextToken](#API_DescribeIndexPolicies_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describeindexpolicies.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describeindexpolicies.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeImportTasks

DescribeLogGroups

All content copied from https://docs.aws.amazon.com/.
