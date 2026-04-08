# DescribeAccountPolicies

Returns a list of all CloudWatch Logs account policies in the account.

To use this operation, you must be signed on with the correct permissions depending on the
type of policy that you are retrieving information for.

- To see data protection policies, you must have the
`logs:GetDataProtectionPolicy` and `logs:DescribeAccountPolicies`
permissions.

- To see subscription filter policies, you must have the
`logs:DescribeSubscriptionFilters` and
`logs:DescribeAccountPolicies` permissions.

- To see transformer policies, you must have the `logs:GetTransformer` and
`logs:DescribeAccountPolicies` permissions.

- To see field index policies, you must have the `logs:DescribeIndexPolicies`
and `logs:DescribeAccountPolicies` permissions.

## Request Syntax

```nohighlight

{
   "accountIdentifiers": [ "string" ],
   "nextToken": "string",
   "policyName": "string",
   "policyType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[accountIdentifiers](#API_DescribeAccountPolicies_RequestSyntax)**

If you are using an account that is set up as a monitoring account for CloudWatch
unified cross-account observability, you can use this to specify the account ID of a source
account. If you do, the operation returns the account policy for the specified account.
Currently, you can specify only one account ID in this parameter.

If you omit this parameter, only the policy in the current account is returned.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 1 item.

Length Constraints: Fixed length of 12.

Pattern: `^\d{12}$`

Required: No

**[nextToken](#API_DescribeAccountPolicies_RequestSyntax)**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[policyName](#API_DescribeAccountPolicies_RequestSyntax)**

Use this parameter to limit the returned policies to only the policy with the name that
you specify.

Type: String

Required: No

**[policyType](#API_DescribeAccountPolicies_RequestSyntax)**

Use this parameter to limit the returned policies to only the policies that match the
policy type that you specify.

Type: String

Valid Values: `DATA_PROTECTION_POLICY | SUBSCRIPTION_FILTER_POLICY | FIELD_INDEX_POLICY | TRANSFORMER_POLICY | METRIC_EXTRACTION_POLICY`

Required: Yes

## Response Syntax

```nohighlight

{
   "accountPolicies": [
      {
         "accountId": "string",
         "lastUpdatedTime": number,
         "policyDocument": "string",
         "policyName": "string",
         "policyType": "string",
         "scope": "string",
         "selectionCriteria": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accountPolicies](#API_DescribeAccountPolicies_ResponseSyntax)**

An array of structures that contain information about the CloudWatch Logs account
policies that match the specified filters.

Type: Array of [AccountPolicy](api-accountpolicy.md) objects

**[nextToken](#API_DescribeAccountPolicies_ResponseSyntax)**

The token to use when requesting the next set of items. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describeaccountpolicies.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describeaccountpolicies.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteTransformer

DescribeConfigurationTemplates
